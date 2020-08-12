package main

import (
	"github.com/wangnengjie/mirai-go"
	"github.com/wangnengjie/mirai-go/model"
	"github.com/wangnengjie/mirai-go/util/json"
	"net/url"
	"strings"
)

func main() {
	c := mirai.NewClient("client1", url.URL{Scheme: "http", Host: "127.0.0.1:8080"}, "12345678")
	b := c.AddBot(123456789, true, 0)
	b.OnStart(getFriendList, getGroupList) // 当前，所有handler的执行是串行的
	b.On(model.GroupMessage, replyFriendList)
	c.Listen(true)
}

func replyFriendList(b *mirai.Bot, m model.MsgRecv) {
	msg, _ := m.(*model.GroupMsg)
	atMe := false
	str := ""
	var mid model.MsgId
	for _, single := range msg.MessageChain {
		switch single.GetType() {
		case model.SourceMsg:
			mid = single.(*model.Source).Id
		case model.AtMsg:
			if single.(*model.At).Target == b.Id() {
				atMe = true
			}
		case model.PlainMsg:
			str += single.(*model.Plain).Text
		}
	}
	if atMe && strings.Contains(str, "好友列表") {
		friendList, _ := b.Data.Load("friendList")
		s, err := json.Json.MarshalIndent(friendList, "", "    ")
		if err != nil {
			b.Log.Errorln(err)
			return
		}
		_, err = b.SendGroupMessage(msg.Sender.Group.Id, model.NewMsgChainBuilder().Plain(string(s)).Done(), mid)
		if err != nil {
			b.Log.Errorln(err)
		}
	}
}

func getFriendList(b *mirai.Bot) {
	if list, err := b.FriendList(); err != nil {
		b.Log.Errorln(err)
	} else {
		b.Log.Debugln(list)
		b.Data.Store("friendList", list)
	}
}

func getGroupList(b *mirai.Bot) {
	if list, err := b.GroupList(); err != nil {
		b.Log.Errorln(err)
	} else {
		b.Log.Debugln(list)
		b.Data.Store("groupList", list)
	}
}
