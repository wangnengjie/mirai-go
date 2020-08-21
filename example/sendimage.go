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
	bot.On(model.FriendMessage, returnImg)
	bot.Loop()
}

func returnImg(ctx *mirai.Context) {
	m := ctx.Message.(*model.FriendMsg)
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
	_, err := ctx.Bot.SendFriendMessage(m.Sender.Id, mc, 0)
	if err != nil {
		ctx.Bot.Log.Error(err)
	}
}
