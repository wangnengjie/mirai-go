package mirai

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/wangnengjie/mirai-go/model"
	"github.com/wangnengjie/mirai-go/util/json"
	"net/url"
	"strconv"
	"sync"
)

type Client struct {
	name        string
	authKey     string
	addr        url.URL
	RestyClient *resty.Client
	Log         *logrus.Entry
	bots        map[model.QQId]*Bot

	debug bool
}

type AuthResp struct {
	Code    int    `json:"code"`
	Session string `json:"session"`
}

type DefaultResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewClient(name string, addr url.URL, authKey string) *Client {
	log := logrus.New()
	c := &Client{
		name:        name,
		authKey:     authKey,
		addr:        addr,
		RestyClient: resty.New(),
		Log:         log.WithFields(logrus.Fields{"Client": name}),
		bots:        make(map[model.QQId]*Bot),
	}
	return c
}

func (c *Client) Listen(debug bool) {
	c.debug = debug
	if debug {
		c.Log.Logger.SetLevel(logrus.DebugLevel)
	}
	c.RestyClient.SetHostURL(c.addr.String()).
		SetHeader("Content-Type", "application/json;charset=utf8")

	wg := sync.WaitGroup{}
	for _, bot := range c.bots {
		session, err := c.auth()
		if err != nil {
			c.Log.Fatalln(err)
		}
		bot.session = session
		err = c.verify(bot, session)
		if err != nil {
			c.Log.Fatalln(err)
		}
		wg.Add(1)
		if c.debug {
			bot.Log.Logger.SetLevel(logrus.DebugLevel)
		}
		c.Log.Infof("Starting [Bot:%d] ...", bot.id)
		go func(b *Bot) {
			b.start()
			wg.Done()
		}(bot)
	}
	wg.Wait()
	c.Log.Infoln("Client exit.")
}

func (c *Client) auth() (string, error) {
	resp, err := c.RestyClient.R().SetBody(map[string]interface{}{"authKey": c.authKey}).Post("/auth")
	if err != nil {
		return "", err
	}
	var r AuthResp
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return "", err
	}
	if r.Code == 1 {
		return "", errors.New("错误的MIRAI API HTTP auth key")
	}
	return r.Session, nil
}

func (c *Client) verify(bot *Bot, session string) error {
	resp, err := c.RestyClient.R().SetBody(map[string]interface{}{"sessionKey": session, "qq": bot.id}).Post("/verify")
	if err != nil {
		return err
	}
	var r DefaultResp
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return err
	}
	return respErrCode(r.Code)
}

func (c *Client) AddBot(id model.QQId, pwd string) *Bot {
	log := logrus.New()
	c.bots[id] = &Bot{
		id:            id,
		pwd:           pwd,
		Log:           log.WithFields(logrus.Fields{"Bot": id}),
		Client:        c,
		session:       "",
		msgHandlers:   make(map[model.MessageRecvType][]func(*Bot, model.MessageRecv)),
		eventHandlers: make(map[model.EventType][]func(*Bot, model.Event)),
	}
	return c.bots[id]
}

func (c *Client) GetBot(id model.QQId) (*Bot, error) {
	b, ok := c.bots[id]
	if ok {
		return b, nil
	}
	return nil, errors.New(fmt.Sprintf("Bot %d does not exist.", id))
}

func respErrCode(code int) error {
	switch code {
	case 0:
		return nil
	case 1:
		return errors.New("错误的auth key")
	case 2:
		return errors.New("指定的Bot不存在")
	case 3:
		return errors.New("session失效或不存在")
	case 4:
		return errors.New("session未认证(未激活)")
	case 5:
		return errors.New("发送消息目标不存在(指定对象不存在)")
	case 6:
		return errors.New("指定文件不存在，出现于发送本地图片")
	case 10:
		return errors.New("无操作权限，指Bot没有对应操作的限权")
	case 20:
		return errors.New("Bot被禁言，指Bot当前无法向指定群发送消息")
	case 30:
		return errors.New("消息过长")
	case 400:
		return errors.New("错误的访问，如参数错误等")
	default:
		return errors.New("unknown error, code:" + strconv.Itoa(code))
	}
}
