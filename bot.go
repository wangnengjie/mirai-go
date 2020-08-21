package mirai

import (
	"errors"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/wangnengjie/mirai-go/model"
	"sync"
	"time"
)

type Mode string

const (
	RecvEventOnly   Mode = "event"
	RecvMessageOnly Mode = "message"
	RecvAll         Mode = "all"
)

type HandlerFunc func(ctx *Context)
type HandlersChan []HandlerFunc

type BotConfig struct {
	Host    string // http://Host[:Port?]
	AuthKey string // authkey for remote mirai client

	Id         model.QQId
	CacheSize  int
	Websocket  bool // 是否使用ws（推荐）
	RecvMode   Mode // M_EVENT, M_MESSAGE, M_EVENT|M_MESSAGE (仅对ws有效)
	FetchMount int  // 使用轮询时每次fetch的信息条数
	Debug      bool // 开启debug模式
	Strict     bool // 严格模式下保证按收到的消息的顺序处理消息 (可能造成性能下降)
}

type Bot struct {
	Mu   sync.RWMutex
	Log  *logrus.Entry
	Data sync.Map

	config     *BotConfig
	sessionKey string
	httpClient *resty.Client
	wsConn     *websocket.Conn
	handlers   map[model.MsgRecvType]HandlersChan
	middleware HandlersChan
	pool       sync.Pool
}

func NewBot(config BotConfig) *Bot {
	bot := &Bot{
		Mu:         sync.RWMutex{},
		Log:        logrus.WithField("Id", config.Id),
		Data:       sync.Map{},
		config:     &config,
		sessionKey: "",
		httpClient: resty.New(),
		wsConn:     nil,
		handlers:   make(map[model.MsgRecvType]HandlersChan),
		middleware: make(HandlersChan, 0),
	}
	bot.httpClient.SetHostURL("http://"+config.Host).
		SetHeader("Content-Type", "application/json;charset=utf8").
		SetLogger(bot.Log)

	bot.httpClient.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		bot.Log.Debugf("%s: %s", req.Method, req.URL)
		return nil
	})
	bot.httpClient.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		bot.Log.Debugf("Response: %s", resp.Status())
		if !resp.IsSuccess() {
			return errors.New("Request to " + resp.Request.URL + " failed")
		}
		return nil
	})

	bot.pool = sync.Pool{New: func() interface{} { return &Context{Bot: bot} }}
	bot.Log.Logger.SetFormatter(&nested.Formatter{
		TimestampFormat: "[2006-01-02 15:04:05]",
		NoColors:        true,
		//CallerFirst:     bot.config.Debug,
	})

	if bot.config.Debug {
		bot.Log.Logger.SetLevel(logrus.DebugLevel)
		bot.Use(func(ctx *Context) {
			ctx.Bot.Log.Debugf("Receive Message: %+v", ctx.Message)
			ctx.Next()
		})
	}
	return bot
}

func (b *Bot) Id() model.QQId {
	return b.config.Id
}

func (b *Bot) SessionKey() string {
	b.Mu.RLock()
	defer b.Mu.RUnlock()
	return b.sessionKey
}

func (b *Bot) SetSessionKey(sessionKey string) {
	b.Mu.Lock()
	b.sessionKey = sessionKey
	b.Mu.Unlock()
}

func (b *Bot) Use(handlers ...HandlerFunc) {
	b.middleware = append(b.middleware, handlers...)
}

func (b *Bot) Connect() error {
	sessionKey, err := b.Auth()
	if err != nil {
		return err
	}
	err = b.Verify(sessionKey)
	if err != nil {
		return err
	}
	b.SetSessionKey(sessionKey)
	err = b.SetConfig(b.config.CacheSize, b.config.Websocket)
	if err != nil {
		return err
	}
	b.Log.Infof("Bot start successfully")
	return nil
}

// 挂载消息和事件监听
func (b *Bot) On(t model.MsgRecvType, handlers ...HandlerFunc) {
	b.handlers[t] = append(b.handlers[t], handlers...)
}

func (b *Bot) Loop() {
	if b.config.Websocket {
		b.wsLoop()
	} else {
		b.fetchLoop()
	}
}

func (b *Bot) wsLoop() {
	u := fmt.Sprintf("ws://%s/%s?sessionKey=%s", b.config.Host, b.config.RecvMode, b.SessionKey())
	wsc, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		b.Log.Panic(err)
	}
	defer wsc.Close()
	b.wsConn = wsc
	b.Log.Info("开始获取消息")
	for {
		_, msgByte, err := b.wsConn.ReadMessage()
		if err != nil {
			b.Log.Panic(err)
		}
		msg, err := model.DeserializeMsgRecv(msgByte)
		if err != nil {
			b.Log.Error(err)
			continue
		}

		if b.config.Strict {
			b.handleMessage(msg)
		} else {
			go b.handleMessage(msg)
		}
	}
}

func (b *Bot) fetchLoop() {
	retry := 0
	b.Log.Info("开始获取消息")
	for {
		msg, err := b.FetchMessage(b.config.FetchMount)
		if err != nil {
			b.Log.Warn(err)
			retry++
			if retry > 10 {
				b.Log.Panic("Fail to fetch message")
			}
			time.Sleep(100 * time.Millisecond)
			continue
		}
		retry = 0
		if b.config.Strict {
			b.handleMessages(msg)
		} else {
			go b.handleMessages(msg)
		}
		if len(msg) < b.config.FetchMount {
			time.Sleep(1 * time.Second)
		}
	}
}

func (b *Bot) handleMessage(message model.MsgRecv) {
	ctx := b.pool.Get().(*Context)
	ctx.reset()
	ctx.Message = message
	ctx.handlers = append(b.middleware, b.handlers[message.GetType()]...)
	ctx.Next()
	b.pool.Put(ctx)
}

func (b *Bot) handleMessages(messages []model.MsgRecv) {
	for _, msg := range messages {
		b.handleMessage(msg)
	}
}
