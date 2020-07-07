package mirai

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/wangnengjie/mirai-go/model"
	"github.com/wangnengjie/mirai-go/util/json"
	"sync"
)

type Bot struct {
	id            model.QQId
	pwd           string
	Log           *logrus.Entry
	Client        *Client
	session       string
	msgHandlers   map[model.MessageRecvType][]func(*Bot, model.MessageRecv)
	eventHandlers map[model.EventType][]func(*Bot, model.Event)
}

func (b *Bot) start() {
	addr := b.Client.addr
	addr.Scheme = "ws"
	addr.Path = "/message"
	addr.RawQuery = "sessionKey=" + b.session
	msgConn, _, err := websocket.DefaultDialer.Dial(addr.String(), nil)
	if err != nil {
		b.Log.Errorln(err)
		return
	}
	defer msgConn.Close()
	addr.Path = "/event"
	eventConn, _, err := websocket.DefaultDialer.Dial(addr.String(), nil)
	if err != nil {
		b.Log.Errorln(err)
		return
	}
	defer eventConn.Close()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		b.messageLoop(msgConn)
		wg.Done()
	}()
	go func() {
		b.eventLoop(eventConn)
		wg.Done()
	}()
	b.Log.Infoln("Bot started successfully.")
	wg.Wait()
}

func (b *Bot) OnMsg(t model.MessageRecvType, handler ...func(bot *Bot, msg model.MessageRecv)) {
	b.msgHandlers[t] = append(b.msgHandlers[t], handler...)
}

func (b *Bot) OnEvent(t model.EventType, handler ...func(bot *Bot, msg model.Event)) {
	b.eventHandlers[t] = append(b.eventHandlers[t], handler...)
}

type SendMsgResp struct {
	DefaultResp
	MessageId model.MessageId `json:"messageId"`
}

func (b *Bot) SendGroupMsg(group model.GroupId, mc model.MsgChain, quoteId *model.MessageId) (*SendMsgResp, error) {
	body := map[string]interface{}{
		"sessionKey":   b.session,
		"group":        group,
		"messageChain": mc,
	}
	if quoteId != nil {
		body["quote"] = *quoteId
	}
	b.Log.Debugln(body)
	resp, err := b.Client.RestyClient.R().SetBody(body).Post("/sendGroupMessage")
	if err != nil {
		return nil, err
	}
	var r SendMsgResp
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return nil, err
	}
	return &r, respErrCode(r.Code)
}
