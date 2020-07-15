package main

import (
	"github.com/wangnengjie/mirai-go"
	"github.com/wangnengjie/mirai-go/model"
	"net/url"
)

func main() {
	c := mirai.NewClient("client1", url.URL{Scheme: "http", Host: "127.0.0.1:8080"}, "12345678")
	// 使用c.Log.Logger.SetOutput()可定向client log输出
	// c.RestyClient可以设置请求的各项内容，hostUrl已被设定为第二个参数的内容
	b := c.AddBot(123456789, true, 0)
	// 使用b.Log.Logger.SetOutput()可定向bot log输出
	b.On(model.GroupMessage, repeat)
	c.Listen(true)
}

func repeat(b *mirai.Bot, msg model.MsgRecv) { // 复读群消息
	m, _ := msg.(*model.GroupMsg)
	// 0 代表不回复消息，msgId是发出的消息的id
	// chain中第一位为source
	msgId, err := b.SendGroupMessage(m.Sender.Group.Id, m.MessageChain[1:], 0)
	// msgId 是刚刚发送的这条消息的id
	if err != nil {
		b.Log.Errorln(err)
	} else {
		b.Log.Debugln(msgId)
	}
}
