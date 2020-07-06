package mirai

import (
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/wangnengjie/mirai-go/model"
	"github.com/wangnengjie/mirai-go/util/json"
	"net/url"
)

type Bot struct {
	id            model.QQId
	pwd           string
	Log           *logrus.Entry
	c             *resty.Client
	session       string
	msgHandlers   map[string]interface{}
	eventHandlers map[string]interface{}
}

func (b *Bot) start(addr url.URL) {
	addr.Scheme = "ws"
	addr.Path = "/message"
	addr.RawQuery = "sessionKey=" + b.session
	b.Log.Infoln(addr.String())
	c, _, err := websocket.DefaultDialer.Dial(addr.String(), nil)
	if err != nil {
		b.Log.Errorln(err)
		return
	}
	defer c.Close()
	b.Log.Infoln("Bot started successfully.")
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			b.Log.Errorln(err)
			return
		}
		//b.Log.Debugln("[Receive Message]:", string(message))

		var msgType model.MessageRecvBase
		err = json.Unmarshal(message, &msgType)
		if err != nil {
			b.Log.Errorln(err)
			continue
		}
		var mc model.MsgChain
		var msg model.MessageRecv
		buf := json.Get(message, "messageChain")
		stream := jsoniter.NewStream(json.Json, nil, buf.Size())
		buf.WriteTo(stream)
		mc, err = model.Deserialize(stream.Buffer())
		if err != nil {
			b.Log.Errorln(err)
			continue
		}

		switch msgType.Type {
		case model.FRIENDMESSAGE:
			var fm model.FriendMsg
			fm.MsgChain = mc
			err = json.Unmarshal(message, &fm)
			msg = &fm
		case model.GROUPMESSAGE:
			var fm model.GroupMsg
			fm.MsgChain = mc
			err = json.Unmarshal(message, &fm)
			msg = &fm
		case model.TEMPMESSAGE:
			var fm model.GroupMsg
			fm.MsgChain = mc
			err = json.Unmarshal(message, &fm)
			msg = &fm
		}
		if err != nil {
			b.Log.Errorln(err)
			continue
		}
		b.Log.Debugln("[Marshal success]:", msg)
	}
}
