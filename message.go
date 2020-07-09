package mirai

import (
	"errors"
	"github.com/wangnengjie/mirai-go/model"
	"github.com/wangnengjie/mirai-go/util/json"
	"strconv"
)

type SendMsgReq struct {
	SessionKey   string         `json:"sessionKey"`
	MessageChain model.MsgChain `json:"messageChain"`
	QQ           model.QQId     `json:"qq,omitempty"`
	Group        model.GroupId  `json:"group,omitempty"`
	Quote        model.MsgId    `json:"quote,omitempty"`
}

type SendImgMsgReq struct {
	SessionKey string        `json:"sessionKey"`
	Urls       []string      `json:"urls"`
	QQ         model.QQId    `json:"qq,omitempty"`
	Group      model.GroupId `json:"group,omitempty"`
}

type RecallReq struct {
	SessionKey string      `json:"sessionKey"`
	Target     model.MsgId `json:"target"`
}

type SendMsgResp struct {
	DefaultResp
	MessageId model.MsgId `json:"messageId"`
}

type ImgUploadResp struct {
	ImageId string `json:"imageId"`
	Url     string `json:"url"`
	Path    string `json:"path"`
}

func (b *Bot) sendMsg(qq model.QQId, group model.GroupId, mc model.MsgChain, quoteId model.MsgId, path string) (model.MsgId, error) {
	body, err := json.Marshal(&SendMsgReq{
		SessionKey:   b.sessionKey,
		MessageChain: mc,
		QQ:           qq,
		Group:        group,
		Quote:        quoteId,
	})
	if err != nil {
		return 0, err
	}
	resp, err := b.Client.RestyClient.R().SetBody(body).Post(path)
	if err != nil {
		return 0, err
	}
	var r SendMsgResp
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return 0, err
	}
	return r.MessageId, respErrCode(r.Code)
}

//向指定群发送消息，quoteId不为0时使用回复
func (b *Bot) SendGroupMessage(group model.GroupId, mc model.MsgChain, quoteId model.MsgId) (model.MsgId, error) {
	return b.sendMsg(0, group, mc, quoteId, "/sendGroupMessage")
}

//向指定好友发送消息，quoteId不为0时使用回复
func (b *Bot) SendFriendMessage(qq model.QQId, mc model.MsgChain, quoteId model.MsgId) (model.MsgId, error) {
	return b.sendMsg(qq, 0, mc, quoteId, "/sendFriendMessage")
}

//向临时会话对象发送消息，quoteId不为0时使用回复
func (b *Bot) SendTempMessage(qq model.QQId, group model.GroupId, mc model.MsgChain, quoteId model.MsgId) (model.MsgId, error) {
	return b.sendMsg(qq, group, mc, quoteId, "/sendTempMessage")
}

//向指定对象（群或好友）发送图片消息（通过url）
//
//除非需要通过此手段获取imageId，否则不推荐使用该接口
//
//当qq和group同时存在时，表示发送临时会话图片（默认为0）
func (b *Bot) SendImageMessage(qq model.QQId, group model.GroupId, urls []string) ([]string, error) {
	body, err := json.Marshal(&SendImgMsgReq{
		SessionKey: b.sessionKey,
		Urls:       urls,
		QQ:         qq,
		Group:      group,
	})
	if err != nil {
		return nil, err
	}
	resp, err := b.Client.RestyClient.R().SetBody(body).Post("/sendImageMessage")
	if err != nil {
		return nil, err
	}
	var imgIds []string
	err = json.Unmarshal(resp.Body(), &imgIds)
	if err != nil {
		return nil, err
	}
	return imgIds, nil
}

//上传图片文件至服务器并返回ImageId
func (b *Bot) UploadImage(t string, path string) (*ImgUploadResp, error) {
	if t != "friend" && t != "group" && t != "temp" {
		return nil, errors.New("type should be friend, group or temp")
	}
	resp, err := b.Client.RestyClient.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetFile("img", path).
		SetFormData(map[string]string{
			"sessionKey": b.sessionKey,
			"type":       t,
		}).Post("/uploadImage")
	if err != nil {
		return nil, err
	}
	var r ImgUploadResp
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

//撤回指定消息。对于bot发送的消息，有2分钟时间限制。对于撤回群聊中群员的消息，需要有相应权限
func (b *Bot) Recall(msgId model.MsgId) error {
	body, err := json.Marshal(&RecallReq{
		SessionKey: b.sessionKey,
		Target:     msgId,
	})
	if err != nil {
		return err
	}
	resp, err := b.Client.RestyClient.R().SetBody(body).Post("/sendImageMessage")
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

func (b *Bot) getMsg(count int, path string) ([]model.MsgRecv, error) {
	if count < 0 {
		return nil, errors.New("count should be larger than 0")
	}
	resp, err := b.Client.RestyClient.R().SetQueryParams(map[string]string{
		"sessionKey": b.sessionKey,
		"count":      strconv.Itoa(count),
	}).Get(path)
	if err != nil {
		return nil, err
	}
	code := json.Get(resp.Body(), "code").ToInt()
	err = respErrCode(code)
	if err != nil {
		return nil, err
	}
	var msgRecvType []model.MsgRecvBase
	buf := []byte(json.Get(resp.Body(), "data").ToString())
	err = json.Unmarshal(buf, &msgRecvType)
	if err != nil {
		return nil, err
	}
	msgs := make([]model.MsgRecv, len(msgRecvType))
	for i := range msgRecvType {
		m, err := model.DeserializeMsgRecv([]byte(json.Get(buf, i).ToString()))
		if err != nil {
			return nil, err
		}
		msgs[i] = m
	}
	return msgs, nil
}

//获取bot接收到的最老消息和最老各类事件(会从MiraiApiHttp消息记录中删除)
func (b *Bot) FetchMessage(count int) ([]model.MsgRecv, error) {
	return b.getMsg(count, "/fetchMessage")
}

//获取bot接收到的最新消息和最新各类事件(会从MiraiApiHttp消息记录中删除)
func (b *Bot) FetchLatestMessage(count int) ([]model.MsgRecv, error) {
	return b.getMsg(count, "/fetchLatestMessage")
}

//获取bot接收到的最老消息和最老各类事件(不会从MiraiApiHttp消息记录中删除)
func (b *Bot) PeekMessage(count int) ([]model.MsgRecv, error) {
	return b.getMsg(count, "/peekMessage")
}

//获取bot接收到的最新消息和最新各类事件(不会从MiraiApiHttp消息记录中删除)
func (b *Bot) PeekLatestMessage(count int) ([]model.MsgRecv, error) {
	return b.getMsg(count, "/peekLatestMessage")
}

//通过messageId获取一条被缓存的消息。当该messageId没有被缓存或缓存失效时，返回code 5(指定对象不存在)
func (b *Bot) MessageFromId(msgId model.MsgId) (model.MsgRecv, error) {
	resp, err := b.Client.RestyClient.R().SetQueryParams(map[string]string{
		"sessionKey": b.sessionKey,
		"id":         strconv.Itoa(int(msgId)),
	}).Get("/messageFromId")
	if err != nil {
		return nil, err
	}
	msg, err := model.DeserializeMsgRecv(resp.Body())
	if err != nil {
		return nil, err
	}
	return msg, nil
}

//获取bot接收并缓存的消息总数，注意不包含被删除的
func (b *Bot) CountMessage() (int, error) {
	resp, err := b.Client.RestyClient.R().SetQueryParam("sessionKey", b.sessionKey).Get("/countMessage")
	if err != nil {
		return 0, err
	}
	code := json.Get(resp.Body(), "code").ToInt()
	data := json.Get(resp.Body(), "data").ToInt()
	return data, respErrCode(code)
}
