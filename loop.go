package mirai

import (
	"github.com/gorilla/websocket"
	"github.com/wangnengjie/mirai-go/model"
)

func (b *Bot) messageLoop(c *websocket.Conn) {
	for {
		_, msgByte, err := c.ReadMessage()
		if err != nil {
			b.Log.Errorln(err)
			return
		}
		go func() {
			msg, err := model.DeserializeMessageRecv(msgByte)
			if err != nil {
				b.Log.Errorln(err)
			}
			b.Log.Debugln(msg)
			handlers, ok := b.msgHandlers[msg.GetType()]
			if !ok || len(handlers) == 0 {
				return
			}
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
			event, err := model.DeserializeEvent(eventByte)
			if err != nil {
				b.Log.Errorln(err)
			}
			handlers, ok := b.eventHandlers[event.GetType()]
			if !ok || len(handlers) == 0 {
				return
			}
			for _, handler := range handlers {
				handler(b, event)
			}
		}()
	}
}
