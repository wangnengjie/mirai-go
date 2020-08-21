package main

import (
	"github.com/wangnengjie/mirai-go"
	"github.com/wangnengjie/mirai-go/model"
)

func main() {
	bot := mirai.NewBot(mirai.BotConfig{
		Host:      "127.0.0.1:8080",
		AuthKey:   "12345678",
		Id:        123456789,
		Websocket: true,
		RecvMode:  mirai.RecvAll,
		Debug:     true,
	})
	err := bot.Connect()
	if err != nil {
		bot.Log.Error(err)
	}
	bot.On(model.GroupMessage, repeat)
	bot.Loop()
}

func repeat(ctx *mirai.Context) { // 复读群消息
	m, _ := ctx.Message.(*model.GroupMsg)
	// 0 代表不回复消息，msgId是发出的消息的id
	// chain中第一位为source
	msgId, err := ctx.Bot.SendGroupMessage(m.Sender.Group.Id, m.MessageChain[1:], 0)
	// msgId 是刚刚发送的这条消息的id
	if err != nil {
		ctx.Bot.Log.Error(err)
	} else {
		ctx.Bot.Log.Info(msgId)
	}
}
