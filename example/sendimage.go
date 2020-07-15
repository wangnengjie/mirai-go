package main

import (
	"github.com/wangnengjie/mirai-go"
	"github.com/wangnengjie/mirai-go/model"
	"net/url"
)

func main() {
	c := mirai.NewClient("client1", url.URL{Scheme: "http", Host: "127.0.0.1:8080"}, "12345678")
	b := c.AddBot(123456789, true, 0)
	b.On(model.FriendMessage, returnImg)
	c.Listen(true)
}

func returnImg(b *mirai.Bot, msg model.MsgRecv) {
	m := msg.(*model.FriendMsg)
	path := ""
	for _, single := range m.MessageChain {
		switch single.GetType() {
		case model.PlainMsg:
			path = single.(*model.Plain).Text
		}
	}
	mc := model.NewMsgChainBuilder().Image("url", path).Done()
	// mc := model.NewMsgChainBuilder().Image("imageId", path).Done()
	// mc := model.NewMsgChainBuilder().Image("path", path).Done() ; Relative to "plugins/MiraiAPIHTTP/images"
	_, err := b.SendFriendMessage(m.Sender.Id, mc, 0)
	if err != nil {
		b.Log.Errorln(err)
	}
}
