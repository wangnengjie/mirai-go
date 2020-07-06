package model

import "github.com/wangnengjie/mirai-go/util/json"

type MessageRecvType string

const (
	GroupMessage  MessageRecvType = "GroupMessage"
	FriendMessage MessageRecvType = "FriendMessage"
	TempMessage   MessageRecvType = "TempMessage"
)

type MessageRecv interface {
	String() string
	GetType() MessageRecvType
}

type MessageRecvBase struct {
	Type MessageRecvType `json:"type"`
}

type GroupMsg struct {
	MessageRecvBase
	MsgChain MsgChain `json:"messageChain"`
	Sender   Member   `json:"sender"`
}

type FriendMsg struct {
	MessageRecvBase
	MsgChain MsgChain `json:"messageChain"`
	Sender   User     `json:"sender"`
}

type TempMsg struct {
	MessageRecvBase
	MsgChain MsgChain `json:"messageChain"`
	Sender   Member   `json:"sender"`
}

func (m *MessageRecvBase) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *MessageRecvBase) GetType() MessageRecvType {
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
