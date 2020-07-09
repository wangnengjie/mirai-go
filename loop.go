package mirai

import (
	"github.com/gorilla/websocket"
	"github.com/wangnengjie/mirai-go/model"
	"time"
)

func (b *Bot) msgLoop(c *websocket.Conn) {
	if c != nil {
		for {
			_, msgByte, err := c.ReadMessage()
			if err != nil {
				close(b.msgCh)
				b.Log.Errorln(err)
				return
			}
			go func() {
				msg, err := model.DeserializeMsgRecv(msgByte)
				b.Log.Debugln(msg)
				if err != nil {
					b.Log.Errorln(err)
					return
				}
				b.msgCh <- msg
			}()
		}
	} else {
		for {
			msg, err := b.FetchMessage(int(b.fetchMount))
			b.Log.Debugln(msg)
			if err != nil {
				b.Log.Errorln(err)
			}
			for _, m := range msg {
				b.msgCh <- m
			}
			if len(msg) < int(b.fetchMount) {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
