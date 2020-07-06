package model

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/wangnengjie/mirai-go/util/json"
)

type MsgChain []Message

type MCBuilder struct {
	mc MsgChain
}

func NewMsgChainBuilder() *MCBuilder {
	return &MCBuilder{}
}

func (mb *MCBuilder) Quote(id MessageId, groupId GroupId, senderId QQId, targetId QQId, origin MsgChain) *MCBuilder {
	mb.mc = append(mb.mc, &Quote{
		MessageBase: MessageBase{QuoteMsg},
		Id:          id,
		GroupId:     groupId,
		SenderId:    senderId,
		TargetId:    targetId,
		Origin:      origin,
	})
	return mb
}

func (mb *MCBuilder) At(targetId QQId) *MCBuilder {
	mb.mc = append(mb.mc, &At{
		MessageBase: MessageBase{AtMsg},
		Target:      targetId,
	})
	return mb
}

func (mb *MCBuilder) AtAll() *MCBuilder {
	mb.mc = append(mb.mc, &AtAll{MessageBase: MessageBase{AtAllMsg}})
	return mb
}

// faceName has lower priority than faceId
//
// faceId: 0 for using faceName; faceName: "" if you don't use it
//
// is there anyone can tell me the Mapping relations of faceId???
func (mb *MCBuilder) Face(faceId int, faceName string) *MCBuilder {
	mb.mc = append(mb.mc, &Face{
		MessageBase: MessageBase{FaceMsg},
		FaceId:      faceId,
		Name:        faceName,
	})
	return mb
}

func (mb *MCBuilder) Plain(text string) *MCBuilder {
	mb.mc = append(mb.mc, &Plain{
		MessageBase: MessageBase{PlainMsg},
		Text:        text,
	})
	return mb
}

// t can be "imageId" "url" "path"
//
// imageId "{xxx}.(mirai)"; url "http(s)://xxx" ; path Relative to "plugins/MiraiAPIHTTP/images"
func (mb *MCBuilder) Image(t string, v string) *MCBuilder {
	m := &Image{}
	m.Type = ImageMsg
	switch t {
	case "imageId":
		m.ImageId = v
	case "url":
		m.Url = v
	case "path":
		m.Path = v
	}
	mb.mc = append(mb.mc, m)
	return mb
}

func (mb *MCBuilder) FlashImage(t string, v string) *MCBuilder {
	m := &FlashImage{}
	m.Type = FlashImageMsg
	switch t {
	case "imageId":
		m.ImageId = v
	case "url":
		m.Url = v
	case "path":
		m.Path = v
	}
	mb.mc = append(mb.mc, m)
	return mb
}

func (mb *MCBuilder) Xml(xml string) *MCBuilder {
	mb.mc = append(mb.mc, &Xml{
		MessageBase: MessageBase{XmlMsg},
		Xml:         xml,
	})
	return mb
}

func (mb *MCBuilder) Json(j string) *MCBuilder {
	mb.mc = append(mb.mc, &Json{
		MessageBase: MessageBase{JsonMsg},
		Json:        j,
	})
	return mb
}

func (mb *MCBuilder) App(content string) *MCBuilder {
	mb.mc = append(mb.mc, &App{
		MessageBase: MessageBase{AppMsg},
		Content:     content,
	})
	return mb
}

func (mb *MCBuilder) Poke(name PokeName) *MCBuilder {
	mb.mc = append(mb.mc, &Poke{
		MessageBase: MessageBase{PokeMsg},
		Name:        name,
	})
	return mb
}

func (mb *MCBuilder) Done() MsgChain {
	return mb.mc
}

func Deserialize(rawjson []byte) (MsgChain, error) {
	var typeList []MessageBase
	err := json.Unmarshal(rawjson, &typeList)
	mc := make(MsgChain, 0, len(typeList))
	if err != nil {
		return mc, err
	}
	for i := range typeList {
		switch typeList[i].Type {
		case SourceMsg:
			mc = append(mc, &Source{})
		case AtMsg:
			mc = append(mc, &At{})
		case AtAllMsg:
			mc = append(mc, &AtAll{})
		case FaceMsg:
			mc = append(mc, &Face{})
		case PlainMsg:
			mc = append(mc, &Plain{})
		case ImageMsg:
			mc = append(mc, &Image{})
		case FlashImageMsg:
			mc = append(mc, &FlashImage{})
		case XmlMsg:
			mc = append(mc, &Xml{})
		case JsonMsg:
			mc = append(mc, &Json{})
		case AppMsg:
			mc = append(mc, &App{})
		case PokeMsg:
			mc = append(mc, &Poke{})
		case QuoteMsg:
			b := json.Get(rawjson, i, "origin")
			stream := jsoniter.NewStream(json.Json, nil, b.Size())
			b.WriteTo(stream)
			subMc, err := Deserialize(stream.Buffer())
			if err != nil {
				return mc, err
			}
			mc = append(mc, &Quote{Origin: subMc})
		case UnknownMsg:
			mc = append(mc, &MessageBase{})
		}
	}
	err = json.Unmarshal(rawjson, &mc)
	if err != nil {
		return mc, err
	}
	return mc, nil
}
