package model

import (
	"github.com/wangnengjie/mirai-go/util/json"
)

type MessageRecvType string
type MessageType string
type MessageId int
type PokeName string
type QQId int64
type GroupId int64

const (
	GROUPMESSAGE  MessageRecvType = "GROUPMESSAGE"
	FRIENDMESSAGE MessageRecvType = "FRIENDMESSAGE"
	TEMPMESSAGE   MessageRecvType = "TempMessage"
)

const (
	SourceMsg     MessageType = "Source"
	QuoteMsg      MessageType = "Quote"
	AtMsg         MessageType = "At"
	AtAllMsg      MessageType = "AtAll"
	FaceMsg       MessageType = "Face"
	PlainMsg      MessageType = "Plain"
	ImageMsg      MessageType = "Image"
	FlashImageMsg MessageType = "FlashImage"
	XmlMsg        MessageType = "Xml"
	JsonMsg       MessageType = "Json"
	AppMsg        MessageType = "App"
	PokeMsg       MessageType = "Poke"
	UnknownMsg    MessageType = "Unknown"
)

const (
	POKE        PokeName = "Poke"
	SHOWLOVE    PokeName = "ShowLove"
	LIKE        PokeName = "Like"
	HEARTBROKEN PokeName = "Heartbroken"
	SIXSIXSIX   PokeName = "SixSixSix"
	FANGDAZHAO  PokeName = "FangDaZhao"
)

type Message interface {
	String() string
}

type MessageBase struct {
	Type MessageType `json:"type"`
}

type Source struct {
	Type MessageType `json:"type"`
	Id   MessageId   `json:"id"`
	Time int         `json:"time"`
}

type Quote struct {
	Type     MessageType `json:"type"`
	Id       MessageId   `json:"id"`
	GroupId  GroupId     `json:"groupId"`
	SenderId QQId        `json:"senderId"`
	TargetId QQId        `json:"targetId"`
	Origin   MsgChain    `json:"origin"`
}

type At struct {
	Type    MessageType `json:"type"`
	Target  QQId        `json:"target"`
	Display string      `json:"display,omitempty"`
}

type AtAll struct {
	Type MessageType `json:"type"`
}

type Face struct {
	Type   MessageType `json:"type"`
	FaceId int         `json:"faceId,omitempty"`
	Name   string      `json:"name,omitempty"`
}

type Plain struct {
	Type MessageType `json:"type"`
	Text string      `json:"text"` // 内容
}

type Image struct {
	Type    MessageType `json:"type"`
	ImageId string      `json:"imageId,omitempty"`
	Url     string      `json:"url,omitempty"`
	Path    string      `json:"path,omitempty"`
}

type FlashImage struct {
	Type    MessageType `json:"type"`
	ImageId string      `json:"imageId,omitempty"`
	Url     string      `json:"url,omitempty"`
	Path    string      `json:"path,omitempty"`
}

type Xml struct {
	Type MessageType `json:"type"`
	Xml  string      `json:"xml"`
}

type Json struct {
	Type MessageType `json:"type"`
	Json string      `json:"json"`
}

type App struct {
	Type    MessageType `json:"type"`
	Content string      `json:"content"`
}

type Poke struct {
	Type MessageType `json:"type"`
	Name PokeName    `json:"name"`
}

type GroupMsg struct {
	Type     MessageRecvType `json:"type"`
	MsgChain MsgChain        `json:"messageChain"`
	Sender   Member          `json:"sender"`
}

type FriendMsg struct {
	Type     MessageRecvType `json:"type"`
	MsgChain MsgChain        `json:"messageChain"`
	Sender   User            `json:"sender"`
}

type TempMsg struct {
	Type     MessageRecvType `json:"type"`
	MsgChain MsgChain        `json:"messageChain"`
	Sender   Member          `json:"sender"`
}

func (m *MessageBase) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *Source) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *Quote) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *At) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *AtAll) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *Face) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *Plain) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *Image) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *FlashImage) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *Xml) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *Json) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *App) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *Poke) String() string {
	s, _ := json.MarshalToString(m)
	return s
}
