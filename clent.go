package mirai

import (
	"errors"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/wangnengjie/mirai-go/model"
	"github.com/wangnengjie/mirai-go/util/json"
	"net/url"
	"strconv"
	"sync"
)

type Client struct {
	name    string
	authKey string
	addr    url.URL
	c       *resty.Client
	Log     *logrus.Entry
	bots    map[model.QQId]*Bot

	debug bool
}

type authResp struct {
	Code    int    `json:"code"`
	Session string `json:"session"`
}

type defaultResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewClient(name, authKey string) *Client {
	log := logrus.New()
	log.SetFormatter(&nested.Formatter{FieldsOrder: []string{"Client"}})
	c := &Client{
		name:    name,
		authKey: authKey,
		c:       nil,
		Log:     log.WithFields(logrus.Fields{"Client": name}),
		bots:    make(map[model.QQId]*Bot),
	}
	return c
}

func (c *Client) Listen(addr url.URL, debug bool) {
	c.addr = addr
	c.debug = debug
	if debug {
		c.Log.Logger.SetLevel(logrus.DebugLevel)
	}
	c.c = resty.New().
		SetHostURL(addr.String()).
		SetHeader("Content-Type", "application/json")

	w := sync.WaitGroup{}
	for _, v := range c.bots {
		session, err := c.auth()
		if err != nil {
			c.Log.Fatalln(err)
		}
		v.session = session
		err = c.verify(v, session)
		if err != nil {
			c.Log.Fatalln(err)
		}
		c.Log.Infoln(session)
		w.Add(1)
		if c.debug {
			v.Log.Logger.SetLevel(logrus.DebugLevel)
		}
		c.Log.Infof("Starting [Bot:%d] ...", v.id)
		go func(b *Bot) {
			b.start(addr)
			w.Done()
		}(v)
	}
	w.Wait()
	c.Log.Infoln("Client exit.")
}

func (c *Client) auth() (string, error) {
	resp, err := c.c.R().SetBody(map[string]interface{}{"authKey": c.authKey}).Post("/auth")
	if err != nil {
		return "", err
	}
	var r authResp
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return "", err
	}
	if r.Code == 1 {
		return "", errors.New("错误的MIRAI API HTTP auth key")
	}
	return r.Session, nil
}

func (c *Client) verify(bot *Bot, session string) error {
	resp, err := c.c.R().SetBody(map[string]interface{}{"sessionKey": session, "qq": bot.id}).Post("/verify")
	if err != nil {
		return err
	}
	var r defaultResp
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return err
	}
	return respErrCode(r.Code)
}

func (c *Client) AddBot(id model.QQId, pwd string) *Bot {
	log := logrus.New()
	log.SetFormatter(&nested.Formatter{FieldsOrder: []string{"Client"}})
	c.bots[id] = &Bot{
		id:            id,
		pwd:           pwd,
		Log:           log.WithFields(logrus.Fields{"Bot": id}),
		c:             c.c,
		session:       "",
		msgHandlers:   nil,
		eventHandlers: nil,
	}
	return c.bots[id]
}

func respErrCode(code int) error {
	switch code {
	case 0:
		return nil
	case 1:
		return errors.New("错误的auth key")
	case 2:
		return errors.New("指定的Bot不存在")
	case 3:
		return errors.New("session失效或不存在")
	case 4:
		return errors.New("session未认证(未激活)")
	case 5:
		return errors.New("发送消息目标不存在(指定对象不存在)")
	case 6:
		return errors.New("指定文件不存在，出现于发送本地图片")
	case 10:
		return errors.New("无操作权限，指Bot没有对应操作的限权")
	case 20:
		return errors.New("Bot被禁言，指Bot当前无法向指定群发送消息")
	case 30:
		return errors.New("消息过长")
	case 400:
		return errors.New("错误的访问，如参数错误等")
	default:
		return errors.New("unknown error, code:" + strconv.Itoa(code))
	}
}
