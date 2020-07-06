package mirai

import (
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/wangnengjie/mirai-go/model"
	"github.com/wangnengjie/mirai-go/util/json"
)

func (b *Bot) messageLoop(c *websocket.Conn) {
	for {
		_, msgByte, err := c.ReadMessage()
		if err != nil {
			b.Log.Errorln(err)
			return
		}
		go func() {
			var msgType model.MessageRecvBase
			err = json.Unmarshal(msgByte, &msgType)
			if err != nil {
				b.Log.Errorln(err)
				return
			}
			handlers, ok := b.msgHandlers[msgType.Type]
			if !ok || len(handlers) == 0 {
				return
			}

			var mc model.MsgChain
			var msg model.MessageRecv
			buf := json.Get(msgByte, "messageChain")
			stream := jsoniter.NewStream(json.Json, nil, buf.Size())
			buf.WriteTo(stream)
			mc, err = model.Deserialize(stream.Buffer())
			if err != nil {
				b.Log.Errorln(err)
				return
			}

			switch msgType.Type {
			case model.FriendMessage:
				var fm model.FriendMsg
				fm.MsgChain = mc
				err = json.Unmarshal(msgByte, &fm)
				msg = &fm
			case model.GroupMessage:
				var fm model.GroupMsg
				fm.MsgChain = mc
				err = json.Unmarshal(msgByte, &fm)
				msg = &fm
			case model.TempMessage:
				var fm model.GroupMsg
				fm.MsgChain = mc
				err = json.Unmarshal(msgByte, &fm)
				msg = &fm
			}
			if err != nil {
				b.Log.Errorln(err)
				return
			}
			b.Log.Debugln("[Marshal success]:", msg)
			for _, handler := range handlers {
				handler(b, msg)
			}
		}()
	}
}

func (b *Bot) eventLoop(c *websocket.Conn) {
	for {
		_, eventByte, err := c.ReadMessage()
		if err != nil {
			b.Log.Errorln(err)
			return
		}
		go func() {
			var eventType model.EventBase
			err := json.Unmarshal(eventByte, &eventType)
			if err != nil {
				b.Log.Errorln(err)
				return
			}
			handlers, ok := b.eventHandlers[eventType.Type]
			if !ok || len(handlers) == 0 {
				return
			}
			var event model.Event
			switch eventType.Type {
			case model.BotOnlineEvent:
				var e model.BotOnline
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.BotOfflineEventActive, model.BotOfflineEventForce, model.BotOfflineEventDropped:
				var e model.BotOffline
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.BotReloginEvent:
				var e model.BotRelogin
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.GroupRecallEvent:
				var e model.GroupRecall
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.FriendRecallEvent:
				var e model.FriendRecall
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.BotGroupPermissionChangeEvent:
				var e model.BotGroupPermissionChange
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.BotMuteEvent:
				var e model.BotMute
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.BotUnmuteEvent:
				var e model.BotUnmute
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.BotJoinGroupEvent:
				var e model.BotJoinGroup
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.BotLeaveEventActive, model.BotLeaveEventKick:
				var e model.BotLeave
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.GroupNameChangeEvent:
				var e model.GroupNameChange
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.GroupEntranceAnnouncementChangeEvent:
				var e model.GroupEntranceAnnouncementChange
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.GroupMuteAllEvent:
				var e model.GroupMuteAll
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.GroupAllowAnonymousChatEvent:
				var e model.GroupAllowAnonymousChat
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.GroupAllowConfessTalkEvent:
				var e model.GroupAllowConfessTalk
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.GroupAllowMemberInviteEvent:
				var e model.GroupAllowMemberInvite
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.MemberJoinEvent:
				var e model.MemberJoin
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.MemberLeaveEventKick, model.MemberLeaveEventQuit:
				var e model.MemberLeave
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.MemberCardChangeEvent:
				var e model.MemberCardChange
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.MemberSpecialTitleChangeEvent:
				var e model.MemberSpecialTitleChange
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.MemberPermissionChangeEvent:
				var e model.MemberPermissionChange
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.MemberMuteEvent:
				var e model.MemberMute
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.MemberUnmuteEvent:
				var e model.MemberUnmute
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.NewFriendRequestEvent:
				var e model.NewFriendRequest
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.MemberJoinRequestEvent:
				var e model.MemberJoinRequest
				err = json.Unmarshal(eventByte, e)
				event = &e
			case model.BotInvitedJoinGroupRequestEvent:
				var e model.BotInvitedJoinGroupRequest
				err = json.Unmarshal(eventByte, e)
				event = &e
			}
			b.Log.Debugln("[Marshal success]:", event)
			for _, handler := range handlers {
				handler(b, event)
			}
		}()
	}
}
