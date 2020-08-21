package mirai

import (
	"errors"
	"github.com/wangnengjie/mirai-go/util"
	"github.com/wangnengjie/mirai-go/util/json"
)

type DefaultResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type AuthResp struct {
	Code    int    `json:"code"`
	Session string `json:"session"`
}

func (b *Bot) Auth() (string, error) {
	resp, err := b.httpClient.R().SetBody(map[string]interface{}{"authKey": b.config.AuthKey}).Post("/auth")
	if err != nil {
		return "", err
	}
	var r AuthResp
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return "", err
	}
	if r.Code == 1 {
		return "", errors.New("错误的MIRAI API HTTP Auth key")
	}
	return r.Session, nil
}

func (b *Bot) Verify(session string) error {
	resp, err := b.httpClient.R().SetBody(map[string]interface{}{"sessionKey": session, "qq": b.Id()}).Post("/verify")
	if err != nil {
		return err
	}
	var r DefaultResp
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return err
	}
	return util.RespErrCode(r.Code)
}

//使用此方式释放session及其相关资源（Bot不会被释放）,不使用的Session应当被释放
//
//长时间（30分钟）未使用的Session将自动释放，否则Session持续保存Bot收到的消息，将会导致内存泄露
//
//开启websocket后将不会自动释放，请务必定期释放
func (b *Bot) Release() error {
	b.Mu.RLock()
	defer b.Mu.RUnlock()
	return b.release()
}

func (b *Bot) release() error {
	resp, err := b.httpClient.R().SetBody(map[string]interface{}{
		"sessionKey": b.sessionKey,
		"qq":         b.Id(),
	}).Post("/release")
	if err != nil {
		return err
	}
	var r DefaultResp
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return err
	}
	return util.RespErrCode(r.Code)
}

func (b *Bot) ReleaseAndReauth() error {
	session, err := b.Auth()
	if err != nil {
		return err
	}
	b.Mu.Lock()
	err = b.release()
	if err != nil {
		return err
	}
	err = b.Verify(session)
	if err != nil {
		return err
	}
	b.sessionKey = session
	b.Mu.Unlock()
	return nil
}
