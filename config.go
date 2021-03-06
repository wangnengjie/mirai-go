package mirai

import (
	"github.com/wangnengjie/mirai-go/util"
	"github.com/wangnengjie/mirai-go/util/json"
)

type Config struct {
	CacheSize       int  `json:"cacheSize"`
	EnableWebsocket bool `json:"enableWebsocket"`
}

func (b *Bot) GetConfig() (*Config, error) {
	resp, err := b.httpClient.R().SetQueryParam("sessionKey", b.SessionKey()).Get("/config")
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(resp.Body(), &config)
	if err != nil {
		return nil, err
	}
	return &config, err
}

func (b *Bot) SetConfig(cacheSize int, enableWebsocket bool) error {
	body := map[string]interface{}{
		"sessionKey":      b.SessionKey(),
		"enableWebsocket": enableWebsocket,
	}
	if cacheSize != 0 {
		body["cacheSize"] = cacheSize
	}
	resp, err := b.httpClient.R().SetBody(body).Post("/config")
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
