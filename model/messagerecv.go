package model

import (
	"github.com/wangnengjie/mirai-go/util/json"
)

type MsgRecvType string

const (
	GroupMessage  MsgRecvType = "GroupMessage"
	FriendMessage MsgRecvType = "FriendMessage"
	TempMessage   MsgRecvType = "TempMessage"
)

type MsgRecv interface {
	String() string
	GetType() MsgRecvType
}

type (
	MsgRecvBase struct {
		Type MsgRecvType `json:"type"`
	}

	GroupMsg struct {
		MsgRecvBase
		MessageChain MsgChain `json:"messageChain"`
		Sender       Member   `json:"sender"`
	}

	FriendMsg struct {
		MsgRecvBase
		MessageChain MsgChain `json:"messageChain"`
		Sender       User     `json:"sender"`
	}

	TempMsg struct {
		MsgRecvBase
		MessageChain MsgChain `json:"messageChain"`
		Sender       Member   `json:"sender"`
	}
)

func (m *MsgRecvBase) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *MsgRecvBase) GetType() MsgRecvType {
	return m.Type
}

func (m *FriendMsg) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *GroupMsg) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *TempMsg) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func DeserializeMsgRecv(rawjson []byte) (MsgRecv, error) {
	var msgType MsgRecvBase
	err := json.Unmarshal(rawjson, &msgType)
	if err != nil {
		return nil, err
	}
	var msg MsgRecv
	switch msgType.Type {
	case FriendMessage, GroupMessage, TempMessage:
		var mc MsgChain
		mc, err = DeserializeMsgChain([]byte(json.Get(rawjson, "messageChain").ToString()))
		if err != nil {
			return nil, err
		}
		switch msgType.Type {
		case FriendMessage:
			var fm FriendMsg
			fm.MessageChain = mc
			err = json.Unmarshal(rawjson, &fm)
			msg = &fm
		case GroupMessage:
			var fm GroupMsg
			fm.MessageChain = mc
			err = json.Unmarshal(rawjson, &fm)
			msg = &fm
		case TempMessage:
			var fm GroupMsg
			fm.MessageChain = mc
			err = json.Unmarshal(rawjson, &fm)
			msg = &fm
		}
	case BotOnlineEvent:
		var e BotOnline
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case BotOfflineEventActive, BotOfflineEventForce, BotOfflineEventDropped:
		var e BotOffline
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case BotReloginEvent:
		var e BotRelogin
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case GroupRecallEvent:
		var e GroupRecall
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case FriendRecallEvent:
		var e FriendRecall
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case BotGroupPermissionChangeEvent:
		var e BotGroupPermissionChange
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case BotMuteEvent:
		var e BotMute
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case BotUnmuteEvent:
		var e BotUnmute
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case BotJoinGroupEvent:
		var e BotJoinGroup
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case BotLeaveEventActive, BotLeaveEventKick:
		var e BotLeave
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case GroupNameChangeEvent:
		var e GroupNameChange
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case GroupEntranceAnnouncementChangeEvent:
		var e GroupEntranceAnnouncementChange
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case GroupMuteAllEvent:
		var e GroupMuteAll
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case GroupAllowAnonymousChatEvent:
		var e GroupAllowAnonymousChat
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case GroupAllowConfessTalkEvent:
		var e GroupAllowConfessTalk
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case GroupAllowMemberInviteEvent:
		var e GroupAllowMemberInvite
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case MemberJoinEvent:
		var e MemberJoin
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case MemberLeaveEventKick, MemberLeaveEventQuit:
		var e MemberLeave
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case MemberCardChangeEvent:
		var e MemberCardChange
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case MemberSpecialTitleChangeEvent:
		var e MemberSpecialTitleChange
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case MemberPermissionChangeEvent:
		var e MemberPermissionChange
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case MemberMuteEvent:
		var e MemberMute
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case MemberUnmuteEvent:
		var e MemberUnmute
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case NewFriendRequestEvent:
		var e NewFriendRequest
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case MemberJoinRequestEvent:
		var e MemberJoinRequest
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	case BotInvitedJoinGroupRequestEvent:
		var e BotInvitedJoinGroupRequest
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	default:
		var e MsgRecvBase
		err = json.Unmarshal(rawjson, &e)
		msg = &e
	}
	if err != nil {
		return nil, err
	}
	return msg, nil
}
