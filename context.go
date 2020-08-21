// Context take example by repo gin-gonic/gin
package mirai

import (
	"github.com/wangnengjie/mirai-go/model"
	"sync"
)

type Context struct {
	Bot     *Bot
	Message model.MsgRecv
	Data    sync.Map

	index    int
	abort    bool
	handlers HandlersChan
}

func (ctx *Context) reset() {
	ctx.Message = nil
	ctx.Data = sync.Map{}
	ctx.index = -1
	ctx.abort = false
	ctx.handlers = nil
}

func (ctx *Context) IsAbort() bool {
	return ctx.abort
}

func (ctx *Context) Abort() {
	ctx.abort = true
}

func (ctx *Context) Next() {
	ctx.index++
	for ctx.index < len(ctx.handlers) && !ctx.IsAbort() {
		ctx.handlers[ctx.index](ctx)
		ctx.index++
	}
}
