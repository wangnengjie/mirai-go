package mirai

import (
	"errors"
	"github.com/wangnengjie/mirai-go/model"
	"github.com/wangnengjie/mirai-go/util/json"
)

type AuthResp struct {
	Code    int    `json:"code"`
	Session string `json:"session"`
}

func (c *Client) auth() (string, error) {
	resp, err := c.RestyClient.R().SetBody(map[string]interface{}{"authKey": c.authKey}).Post("/auth")
	if err != nil {
		return "", err
	}
	var r AuthResp
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return "", err
	}
	if r.Code == 1 {
		return "", errors.New("错误的MIRAI API HTTP auth key")
	}
	return r.Session, nil
}

func (c *Client) verify(qq model.QQId, session string) error {
	resp, err := c.RestyClient.R().SetBody(map[string]interface{}{"sessionKey": session, "qq": qq}).Post("/verify")
	if err != nil {
		return err
	}
	var r DefaultResp
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return err
	}
	return respErrCode(r.Code)
}

//使用此方式释放session及其相关资源（Bot不会被释放）,不使用的Session应当被释放
//
//长时间（30分钟）未使用的Session将自动释放，否则Session持续保存Bot收到的消息，将会导致内存泄露
//
//开启websocket后将不会自动释放，请务必定期释放
func (c *Client) Release(bot *Bot) error {
	bot.Mu.RLock()
	defer bot.Mu.RUnlock()
	return c.release(bot)
}

func (c *Client) release(bot *Bot) error {
	resp, err := c.RestyClient.R().SetBody(map[string]interface{}{
		"sessionKey": bot.sessionKey,
		"qq":         bot.id,
	}).Post("/release")
	if err != nil {
		return err
	}
	var r DefaultResp
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return err
	}
	return respErrCode(r.Code)
}

func (c *Client) ReleaseAndReauth(bot *Bot) error {
	session, err := c.auth()
	if err != nil {
		return err
	}
	bot.Mu.Lock()
	err = c.release(bot)
	if err != nil {
		return err
	}
	err = c.verify(bot.Id(), session)
	if err != nil {
		return err
	}
	bot.sessionKey = session
	bot.Mu.Unlock()
	return nil
}
