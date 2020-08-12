package main

import (
	"fmt"
	"github.com/wangnengjie/mirai-go"
	"net/url"
	"time"
)

func main() {
	c := mirai.NewClient("client1", url.URL{Scheme: "http", Host: "127.0.0.1:8080"}, "12345678")
	b := c.AddBot(123456789, true, 0)
	go testReleaseAndReauth(b)
	c.Listen(true)
}

func testReleaseAndReauth(bot *mirai.Bot) {
	time.Sleep(20 * time.Second)
	fmt.Println("Prev:", bot.SessionKey())
	err := bot.Client.ReleaseAndReauth(bot)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("After:", bot.SessionKey())
}
