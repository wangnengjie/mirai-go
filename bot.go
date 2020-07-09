package mirai

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/wangnengjie/mirai-go/model"
	"sync"
)

type Bot struct {
	Mu sync.RWMutex

	id              model.QQId
	sessionKey      string
	enableWebsocket bool
	fetchMount      uint // 使用轮询时每次fetch的信息条数
	Log             *logrus.Entry
	Client          *Client

	msgCh       chan model.MsgRecv
	msgHandlers map[model.MsgRecvType][]func(*Bot, model.MsgRecv)
	Data        map[string]interface{} // 线程不安全，你可以在里面放置任何东西
}

func (b *Bot) start() {
	err := b.SetConfig(0, b.enableWebsocket)
	if err != nil {
		b.Log.Errorln("Fail to set websocket config")
		return
	}
	var c *websocket.Conn = nil
	if b.enableWebsocket {
		addr := b.Client.addr
		addr.Scheme = "ws"
		addr.Path = "/all"
		addr.RawQuery = "sessionKey=" + b.sessionKey
		c, _, err = websocket.DefaultDialer.Dial(addr.String(), nil)
		if err != nil {
			b.Log.Errorln(err)
			return
		}
		defer c.Close()
	}
	b.Log.Infoln("Bot started successfully.")
	go b.msgLoop(c)
	for msg := range b.msgCh {
		handlers, ok := b.msgHandlers[msg.GetType()]
		if ok && len(handlers) > 0 {
			for _, handler := range handlers {
				handler(b, msg) // 当前handlers同步执行，后续可能改成goroutin
			}
		}
	}
}

func (b *Bot) On(t model.MsgRecvType, handlers ...func(bot *Bot, msg model.MsgRecv)) {
	b.msgHandlers[t] = append(b.msgHandlers[t], handlers...)
}
