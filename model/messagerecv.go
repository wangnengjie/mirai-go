package model

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/wangnengjie/mirai-go/util/json"
)

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

func DeserializeMessageRecv(rawjson []byte) (MessageRecv, error) {
	var msgType MessageRecvBase
	err := json.Unmarshal(rawjson, &msgType)
	if err != nil {
		return nil, err
	}

	var mc MsgChain
	var msg MessageRecv
	buf := json.Get(rawjson, "messageChain")
	stream := jsoniter.NewStream(json.Json, nil, buf.Size())
	buf.WriteTo(stream)
	mc, err = DeserializeMessageChain(stream.Buffer())
	if err != nil {
		return nil, err
	}

	switch msgType.Type {
	case FriendMessage:
		var fm FriendMsg
		fm.MsgChain = mc
		err = json.Unmarshal(rawjson, &fm)
		msg = &fm
	case GroupMessage:
		var fm GroupMsg
		fm.MsgChain = mc
		err = json.Unmarshal(rawjson, &fm)
		msg = &fm
	case TempMessage:
		var fm GroupMsg
		fm.MsgChain = mc
		err = json.Unmarshal(rawjson, &fm)
		msg = &fm
	default:
		err = errors.New("unknown message receive type")
	}
	if err != nil {
		return nil, err
	}
	return msg, nil
}
