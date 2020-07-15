# mirai-go
基于[mirai-api-http](https://github.com/project-mirai/mirai-api-http)的golang sdk

项目目前仍在开发中，所有功能尚未经过测试且可能出现break change，**非常不建议**在生产环境中使用

~~只是个玩具啦~~

目前基于`mirai`的相当一部分工具都处于不稳定状态，大规模使用或者商业使用建议使用[coolq](https://cqp.cc/)

### 获取

目前尚未发版

## Quick Start

```go
package main

import (
	"github.com/wangnengjie/mirai-go"
	"github.com/wangnengjie/mirai-go/model"
	"net/url"
)

func main() {
	c := mirai.NewClient("client1", url.URL{Scheme: "http", Host: "127.0.0.1:8080"}, "12345678")
	// 使用c.Log.Logger.SetOutput()可定向client log输出
	// c.RestyClient可以设置请求的各项内容，hostUrl已被设定为第二个参数的内容
	b := c.AddBot(123456789, true, 0)
	// 使用b.Log.Logger.SetOutput()可定向bot log输出
	b.On(model.GroupMessage, repeat)
	c.Listen(true)
}

func repeat(b *mirai.Bot, msg model.MsgRecv) { // 复读群消息
	m, _ := msg.(*model.GroupMsg)
	// 0 代表不回复消息，msgId是发出的消息的id
	// chain中第一位为source
	msgId, err := b.SendGroupMessage(m.Sender.Group.Id, m.MessageChain[1:], 0)
	// msgId 是刚刚发送的这条消息的id
	if err != nil {
		b.Log.Errorln(err)
	} else {
		b.Log.Debugln(msgId)
	}
}

```

## Todos

- [ ] 添加更多example
- [ ] 添加更多调试信息（不清楚需要添加哪些，欢迎提建议）
- [ ] 完善文档
- [ ] 中间件功能
- [ ] 拦截器功能
- [ ] command接口
- [ ] 测试（目前没有可供测试的条件，qq号申请越来越严了）
- [ ] 性能优化？

go语言刚入门菜鸡，项目可能会出现各种问题，欢迎提issue

## 依赖

- [resty](https://github.com/go-resty/resty)：Simple HTTP and REST client library for Go
- [websocket](https://github.com/gorilla/websocket)：A fast, well-tested and widely used WebSocket implementation for Go
- [jsoniter](https://github.com/json-iterator/go)：A high-performance 100% compatible drop-in replacement of "encoding/json"
- [logrus](https://github.com/sirupsen/logrus)：Structured, pluggable logging for Go

## 鸣谢

特别感谢`mirai`项目组[mamoe](https://github.com/mamoe)

- [mirai](https://github.com/mamoe/mirai)：全开源 高效率 QQ机器人/Android QQ协议支持库 for JVM / Android
- [mirai-console](https://github.com/mamoe/mirai-console)：mirai 的高效率 QQ 机器人控制台
- [mirai-api-http](https://github.com/mamoe/mirai-api-http)：Mirai HTTP API (console) plugin

## 许可证

[GNU AGPLv3](https://choosealicense.com/licenses/agpl-3.0/)，基于`mirai`的一系列项目均使用`GNU AGPLv3`开源许可证，使用时请遵守相关规则