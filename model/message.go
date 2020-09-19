package model

import (
	"github.com/wangnengjie/mirai-go/util/json"
)

type MsgType string
type MsgId int
type PokeName string
type QQId int64
type GroupId int64

const (
	SourceMsg     MsgType = "Source"
	QuoteMsg      MsgType = "Quote"
	AtMsg         MsgType = "At"
	AtAllMsg      MsgType = "AtAll"
	FaceMsg       MsgType = "Face"
	PlainMsg      MsgType = "Plain"
	ImageMsg      MsgType = "Image"
	FlashImageMsg MsgType = "FlashImage"
	VoiceMsg      MsgType = "Voice"
	XmlMsg        MsgType = "Xml"
	JsonMsg       MsgType = "Json"
	AppMsg        MsgType = "App"
	PokeMsg       MsgType = "Poke"
)

const (
	POKE        PokeName = "Poke"
	SHOWLOVE    PokeName = "ShowLove"
	LIKE        PokeName = "Like"
	HEARTBROKEN PokeName = "Heartbroken"
	SIXSIXSIX   PokeName = "SixSixSix"
	FANGDAZHAO  PokeName = "FangDaZhao"
)

type Msg interface {
	String() string
	GetType() MsgType
}

type (
	MsgBase struct {
		Type MsgType `json:"type"`
	}

	Source struct {
		MsgBase
		Id   MsgId `json:"id"`
		Time int   `json:"time"`
	}

	Quote struct {
		MsgBase
		Id       MsgId    `json:"id"`
		GroupId  GroupId  `json:"groupId"`
		SenderId QQId     `json:"senderId"`
		TargetId QQId     `json:"targetId"`
		Origin   MsgChain `json:"origin"`
	}

	At struct {
		MsgBase
		Target  QQId   `json:"target"`
		Display string `json:"display,omitempty"`
	}

	AtAll struct {
		MsgBase
	}

	Face struct {
		MsgBase
		FaceId int    `json:"faceId,omitempty"`
		Name   string `json:"name,omitempty"`
	}

	Plain struct {
		MsgBase
		Text string `json:"text"`
	}

	Image struct {
		MsgBase
		ImageId string `json:"imageId,omitempty"`
		Url     string `json:"url,omitempty"`
		Path    string `json:"path,omitempty"`
	}

	FlashImage struct {
		MsgBase
		ImageId string `json:"imageId,omitempty"`
		Url     string `json:"url,omitempty"`
		Path    string `json:"path,omitempty"`
	}

	Voice struct {
		MsgBase
		VoiceId string `json:"voiceId"`
		Url     string `json:"url"`
		Path    string `json:"path"`
	}

	Xml struct {
		MsgBase
		Xml string `json:"xml"`
	}

	Json struct {
		MsgBase
		Json string `json:"json"`
	}

	App struct {
		MsgBase
		Content string `json:"content"`
	}

	Poke struct {
		MsgBase
		Name PokeName `json:"name"`
	}
)

func (m *MsgBase) GetType() MsgType {
	return m.Type
}

func (m *MsgBase) String() string {
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

func (m *Voice) String() string {
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
