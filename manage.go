package mirai

import (
	"github.com/wangnengjie/mirai-go/model"
	"github.com/wangnengjie/mirai-go/util/json"
	"strconv"
)

type MuteOrUnmuteReq struct {
	SessionKey string        `json:"sessionKey"`
	Target     model.GroupId `json:"target"`
	MemberId   model.QQId    `json:"memberId,omitempty"`
	Time       int           `json:"time,omitempty"`
}

type KickReq struct {
	SessionKey string        `json:"sessionKey"`
	Target     model.GroupId `json:"target"`
	MemberId   model.QQId    `json:"memberId"`
	Msg        string        `json:"msg,omitempty"`
}

type QuitReq struct {
	SessionKey string        `json:"sessionKey"`
	Target     model.GroupId `json:"target"`
}

type SetGroupConfigReq struct {
	SessionKey string            `json:"sessionKey"`
	Target     model.GroupId     `json:"target"`
	Config     model.GroupConfig `json:"config"`
}

type SetMemberInfoReq struct {
	SessionKey string           `json:"sessionKey"`
	Target     model.GroupId    `json:"target"`
	MemberId   model.QQId       `json:"memberId"`
	Info       model.MemberInfo `json:"info"`
}

type DefaultEventReply struct {
	SessionKey string        `json:"sessionKey"`
	EventId    model.EventId `json:"eventId"`
	FromId     model.QQId    `json:"fromId"`
	GroupId    model.GroupId `json:"groupId"`
	Operate    int           `json:"operate"`
	Message    string        `json:"message"`
}

type NewFriendReply DefaultEventReply

type MemberJoinReply DefaultEventReply

type BotInvitedJoinGroupReply DefaultEventReply

//获取bot的好友列表
func (b *Bot) FriendList() ([]model.User, error) {
	resp, err := b.Client.RestyClient.R().SetQueryParam("sessionKey", b.sessionKey).Get("/friendList")
	if err != nil {
		return nil, err
	}
	var u []model.User
	err = json.Unmarshal(resp.Body(), &u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

//获取bot的群列表
func (b *Bot) GroupList() ([]model.Group, error) {
	resp, err := b.Client.RestyClient.R().SetQueryParam("sessionKey", b.sessionKey).Get("/groupList")
	if err != nil {
		return nil, err
	}
	var g []model.Group
	err = json.Unmarshal(resp.Body(), &g)
	if err != nil {
		return nil, err
	}
	return g, nil
}

//获取bot指定群种的成员列表
func (b *Bot) MemberList() ([]model.Member, error) {
	resp, err := b.Client.RestyClient.R().SetQueryParam("sessionKey", b.sessionKey).Get("/memberList")
	if err != nil {
		return nil, err
	}
	var members []model.Member
	err = json.Unmarshal(resp.Body(), &members)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (b *Bot) muteOrUnmute(group model.GroupId, qq model.QQId, time int, path string) error {
	body, err := json.Marshal(&MuteOrUnmuteReq{
		SessionKey: b.sessionKey,
		Target:     group,
		MemberId:   qq,
		Time:       time,
	})
	if err != nil {
		return err
	}
	resp, err := b.Client.RestyClient.R().SetBody(body).Post(path)
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

//令指定群进行全体禁言（需要有相关限权）
func (b *Bot) MuteAll(group model.GroupId) error {
	return b.muteOrUnmute(group, 0, 0, "/muteAll")
}

//令指定群解除全体禁言（需要有相关限权）
func (b *Bot) UnmuteAll(group model.GroupId) error {
	return b.muteOrUnmute(group, 0, 0, "/unmuteAll")
}

//指定群禁言指定群员（需要有相关限权）
func (b *Bot) Mute(group model.GroupId, qq model.QQId, time int) error {
	return b.muteOrUnmute(group, qq, time, "/mute")
}

//指定群解除群成员禁言（需要有相关限权）
func (b *Bot) Unmute(group model.GroupId, qq model.QQId) error {
	return b.muteOrUnmute(group, qq, 0, "/unmute")
}

//移除指定群成员（需要有相关限权）
func (b *Bot) Kick(group model.GroupId, qq model.QQId, msg string) error {
	body, err := json.Marshal(&KickReq{
		SessionKey: b.sessionKey,
		Target:     group,
		MemberId:   qq,
		Msg:        msg,
	})
	if err != nil {
		return err
	}
	resp, err := b.Client.RestyClient.R().SetBody(body).Post("/kick")
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

//使Bot退出群聊
func (b *Bot) Quit(group model.GroupId) error {
	body, err := json.Marshal(&QuitReq{
		SessionKey: b.sessionKey,
		Target:     group,
	})
	if err != nil {
		return err
	}
	resp, err := b.Client.RestyClient.R().SetBody(body).Post("/quit")
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

//修改群设置（需要有相关限权）
func (b *Bot) SetGroupConfig(group model.GroupId, config model.GroupConfig) error {
	body, err := json.Marshal(&SetGroupConfigReq{
		SessionKey: b.sessionKey,
		Target:     group,
		Config:     config,
	})
	if err != nil {
		return err
	}
	resp, err := b.Client.RestyClient.R().SetBody(body).Post("/groupConfig")
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

//获取群设置
func (b *Bot) GetGroupConfig(group model.GroupId) (*model.GroupConfig, error) {
	resp, err := b.Client.RestyClient.R().SetQueryParams(map[string]string{
		"sessionKey": b.sessionKey,
		"target":     strconv.FormatInt(int64(group), 10),
	}).Get("/groupConfig")
	if err != nil {
		return nil, err
	}
	var config model.GroupConfig
	err = json.Unmarshal(resp.Body(), &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

//修改群员资料（需要有相关限权）
func (b *Bot) SetMemberInfo(group model.GroupId, qq model.QQId, info model.MemberInfo) error {
	body, err := json.Marshal(&SetMemberInfoReq{
		SessionKey: b.sessionKey,
		Target:     group,
		MemberId:   qq,
		Info:       info,
	})
	if err != nil {
		return err
	}
	resp, err := b.Client.RestyClient.R().SetBody(body).Post("/memberInfo")
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

//获取群员资料
func (b *Bot) GetMemberInfo(group model.GroupId, qq model.QQId) (*model.MemberInfo, error) {
	resp, err := b.Client.RestyClient.R().SetQueryParams(map[string]string{
		"sessionKey": b.sessionKey,
		"target":     strconv.FormatInt(int64(group), 10),
		"memberId":   strconv.FormatInt(int64(qq), 10),
	}).Get("/memberInfo")
	if err != nil {
		return nil, err
	}
	var info model.MemberInfo
	err = json.Unmarshal(resp.Body(), &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

//回复好友申请
//
//0 同意添加好友
//
//1 拒绝添加好友
//
//2 拒绝添加好友并添加黑名单，不再接收该用户的好友申请
func (b *Bot) NewFriendReply(operate int, msg string, req *model.NewFriendRequest) error {
	body, err := json.Marshal(&NewFriendReply{
		SessionKey: b.sessionKey,
		EventId:    req.EventId,
		FromId:     req.FromId,
		GroupId:    req.GroupId,
		Operate:    operate,
		Message:    msg,
	})
	if err != nil {
		return err
	}
	resp, err := b.Client.RestyClient.R().SetBody(body).Post("/resp/newFriendRequestEvent")
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

//响应用户入群申请（Bot需要有管理员权限）
//
//0 同意入群
//
//1 拒绝入群
//
//2 忽略请求
//
//3 拒绝入群并添加黑名单，不再接收该用户的入群申请
//
//4 忽略入群并添加黑名单，不再接收该用户的入群申请
func (b *Bot) MemberJoinReply(operate int, msg string, req *model.MemberJoinRequest) error {
	body, err := json.Marshal(&MemberJoinReply{
		SessionKey: b.sessionKey,
		EventId:    req.EventId,
		FromId:     req.FromId,
		GroupId:    req.GroupId,
		Operate:    operate,
		Message:    msg,
	})
	if err != nil {
		return err
	}
	resp, err := b.Client.RestyClient.R().SetBody(body).Post("/resp/memberJoinRequestEvent")
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

//响应Bot被邀请入群
//
//0 同意邀请
//
//1 拒绝邀请
func (b *Bot) BotInvitedJoinGroupReply(operate int, msg string, req *model.BotInvitedJoinGroupRequest) error {
	body, err := json.Marshal(&BotInvitedJoinGroupReply{
		SessionKey: b.sessionKey,
		EventId:    req.EventId,
		FromId:     req.FromId,
		GroupId:    req.GroupId,
		Operate:    operate,
		Message:    msg,
	})
	if err != nil {
		return err
	}
	resp, err := b.Client.RestyClient.R().SetBody(body).Post("/resp/botInvitedJoinGroupRequestEvent")
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
