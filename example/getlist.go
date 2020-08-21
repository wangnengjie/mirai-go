package main

import (
	"github.com/wangnengjie/mirai-go"
	"github.com/wangnengjie/mirai-go/model"
	"github.com/wangnengjie/mirai-go/util/json"
	"strings"
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
	getFriendList(bot)
	getGroupList(bot)
	bot.On(model.GroupMessage, replyFriendList)
	bot.Loop()
}

func replyFriendList(ctx *mirai.Context) {
	msg, _ := ctx.Message.(*model.GroupMsg)
	atMe := false
	str := ""
	var mid model.MsgId
	for _, single := range msg.MessageChain {
		switch single.GetType() {
		case model.SourceMsg:
			mid = single.(*model.Source).Id
		case model.AtMsg:
			if single.(*model.At).Target == ctx.Bot.Id() {
				atMe = true
			}
		case model.PlainMsg:
			str += single.(*model.Plain).Text
		}
	}
	if atMe && strings.Contains(str, "好友列表") {
		friendList, _ := ctx.Bot.Data.Load("friendList")
		s, err := json.Json.MarshalIndent(friendList, "", "    ")
		if err != nil {
			ctx.Bot.Log.Error(err)
			return
		}
		_, err = ctx.Bot.SendGroupMessage(msg.Sender.Group.Id, model.NewMsgChainBuilder().Plain(string(s)).Done(), mid)
		if err != nil {
			ctx.Bot.Log.Error(err)
		}
	}
}

func getFriendList(b *mirai.Bot) {
	if list, err := b.FriendList(); err != nil {
		b.Log.Error(err)
	} else {
		b.Log.Debug(list)
		b.Data.Store("friendList", list)
	}
}

func getGroupList(b *mirai.Bot) {
	if list, err := b.GroupList(); err != nil {
		b.Log.Error(err)
	} else {
		b.Log.Debug(list)
		b.Data.Store("groupList", list)
	}
}
