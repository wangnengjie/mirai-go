package model

import (
	"github.com/wangnengjie/mirai-go/util/json"
)

type MessageType string
type MessageId int
type PokeName string
type QQId int64
type GroupId int64

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
	GetType() MessageType
}

type MessageBase struct {
	Type MessageType `json:"type"`
}

type Source struct {
	MessageBase
	Id   MessageId `json:"id"`
	Time int       `json:"time"`
}

type Quote struct {
	MessageBase
	Id       MessageId `json:"id"`
	GroupId  GroupId   `json:"groupId"`
	SenderId QQId      `json:"senderId"`
	TargetId QQId      `json:"targetId"`
	Origin   MsgChain  `json:"origin"`
}

type At struct {
	MessageBase
	Target  QQId   `json:"target"`
	Display string `json:"display,omitempty"`
}

type AtAll struct {
	MessageBase
}

type Face struct {
	MessageBase
	FaceId int    `json:"faceId,omitempty"`
	Name   string `json:"name,omitempty"`
}

type Plain struct {
	MessageBase
	Text string `json:"text"`
}

type Image struct {
	MessageBase
	ImageId string `json:"imageId,omitempty"`
	Url     string `json:"url,omitempty"`
	Path    string `json:"path,omitempty"`
}

type FlashImage struct {
	MessageBase
	ImageId string `json:"imageId,omitempty"`
	Url     string `json:"url,omitempty"`
	Path    string `json:"path,omitempty"`
}

type Xml struct {
	MessageBase
	Xml string `json:"xml"`
}

type Json struct {
	MessageBase
	Json string `json:"json"`
}

type App struct {
	MessageBase
	Content string `json:"content"`
}

type Poke struct {
	MessageBase
	Name PokeName `json:"name"`
}

func (m *MessageBase) GetType() MessageType {
	return m.Type
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
