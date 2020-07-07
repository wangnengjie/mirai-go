package model

import (
	"errors"
	"github.com/wangnengjie/mirai-go/util/json"
)

type EventType string
type EventId int64

const (
	BotOnlineEvent                       EventType = "BotOnlineEvent"                       //Bot登录成功
	BotOfflineEventActive                EventType = "BotOfflineEventActive"                //Bot主动离线
	BotOfflineEventForce                 EventType = "BotOfflineEventForce"                 //Bot被挤下线
	BotOfflineEventDropped               EventType = "BotOfflineEventDropped"               //Bot被服务器断开或因网络问题而掉线
	BotReloginEvent                      EventType = "BotReloginEvent"                      //Bot主动重新登录
	GroupRecallEvent                     EventType = "GroupRecallEvent"                     //群消息撤回
	FriendRecallEvent                    EventType = "FriendRecallEvent"                    //好友消息撤回
	BotGroupPermissionChangeEvent        EventType = "BotGroupPermissionChangeEvent"        //Bot在群里的权限被改变. 操作人一定是群主
	BotMuteEvent                         EventType = "BotMuteEvent"                         //Bot被禁言
	BotUnmuteEvent                       EventType = "BotUnmuteEvent"                       //Bot被取消禁言
	BotJoinGroupEvent                    EventType = "BotJoinGroupEvent"                    //Bot加入了一个新群
	BotLeaveEventActive                  EventType = "BotLeaveEventActive"                  //Bot主动退出一个群
	BotLeaveEventKick                    EventType = "BotLeaveEventKick"                    //Bot被踢出一个群
	GroupNameChangeEvent                 EventType = "GroupNameChangeEvent"                 //某个群名改变
	GroupEntranceAnnouncementChangeEvent EventType = "GroupEntranceAnnouncementChangeEvent" //某群入群公告改变
	GroupMuteAllEvent                    EventType = "GroupMuteAllEvent"                    //全员禁言
	GroupAllowAnonymousChatEvent         EventType = "GroupAllowAnonymousChatEvent"         //匿名聊天
	GroupAllowConfessTalkEvent           EventType = "GroupAllowConfessTalkEvent"           //坦白说
	GroupAllowMemberInviteEvent          EventType = "GroupAllowMemberInviteEvent"          //允许群员邀请好友加群
	MemberJoinEvent                      EventType = "MemberJoinEvent"                      //新人入群的事件
	MemberLeaveEventKick                 EventType = "MemberLeaveEventKick"                 //成员被踢出群（该成员不是Bot）
	MemberLeaveEventQuit                 EventType = "MemberLeaveEventQuit"                 //成员主动离群（该成员不是Bot）
	MemberCardChangeEvent                EventType = "MemberCardChangeEvent"                //群名片改动
	MemberSpecialTitleChangeEvent        EventType = "MemberSpecialTitleChangeEvent"        //群头衔改动（只有群主有操作限权）
	MemberPermissionChangeEvent          EventType = "MemberPermissionChangeEvent"          //成员权限改变的事件（该成员不可能是Bot，见BotGroupPermissionChangeEvent）
	MemberMuteEvent                      EventType = "MemberMuteEvent"                      //群成员被禁言事件（该成员不可能是Bot，见BotMuteEvent）
	MemberUnmuteEvent                    EventType = "MemberUnmuteEvent"                    //群成员被取消禁言事件（该成员不可能是Bot，见BotUnmuteEvent）
	NewFriendRequestEvent                EventType = "NewFriendRequestEvent"                //添加好友申请
	MemberJoinRequestEvent               EventType = "MemberJoinRequestEvent"               //用户入群申请（Bot需要有管理员权限）
	BotInvitedJoinGroupRequestEvent      EventType = "BotInvitedJoinGroupRequestEvent"      //Bot被邀请入群申请
)

type Event interface {
	GetType() EventType
	String() string
}

type EventBase struct {
	Type EventType `json:"type"`
}

func (e *EventBase) GetType() EventType {
	return e.Type
}

func (e *EventBase) String() string {
	s, _ := json.MarshalToString(e)
	return s
}

type BotOnline struct {
	EventBase
	QQ QQId `json:"qq"`
}

type BotOffline struct {
	EventBase
	QQ QQId `json:"qq"`
}

type BotRelogin struct {
	EventBase
	QQ QQId `json:"qq"`
}

type GroupRecall struct {
	EventBase
	AuthorId  QQId      `json:"authorId"`
	MessageId MessageId `json:"messageId"`
	Time      int       `json:"time"`
	Group     Group     `json:"group"`
	Operator  Member    `json:"operator"`
}

type FriendRecall struct {
	EventBase
	AuthorId  QQId      `json:"authorId"`
	MessageId MessageId `json:"messageId"`
	Time      int       `json:"time"`
	Operator  QQId      `json:"operator"`
}

type BotGroupPermissionChange struct {
	EventBase
	Origin  GroupPermission `json:"origin"`
	New     GroupPermission `json:"new"`
	Current GroupPermission `json:"current"`
	Group   Group           `json:"group"`
}

type BotMute struct {
	EventBase
	DurationSeconds int    `json:"durationSeconds"`
	Operator        Member `json:"operator"`
}

type BotUnmute struct {
	EventBase
	Operator Member `json:"operator"`
}

type BotJoinGroup struct {
	EventBase
	Group Group `json:"group"`
}

type BotLeave struct {
	EventBase
	Group Group `json:"group"`
}

type GroupNameChange struct {
	EventBase
	Origin   string `json:"origin"`
	New      string `json:"new"`
	Current  string `json:"current"`
	Group    Group  `json:"group"`
	Operator Member `json:"operator"`
}

type GroupEntranceAnnouncementChange struct {
	EventBase
	Origin   string `json:"origin"`
	New      string `json:"new"`
	Current  string `json:"current"`
	Group    Group  `json:"group"`
	Operator Member `json:"operator"`
}

type GroupMuteAll struct {
	EventBase
	Origin   bool   `json:"origin"`
	New      bool   `json:"new"`
	Current  bool   `json:"current"`
	Group    Group  `json:"group"`
	Operator Member `json:"operator"`
}

type GroupAllowAnonymousChat struct {
	EventBase
	Origin   bool   `json:"origin"`
	New      bool   `json:"new"`
	Current  bool   `json:"current"`
	Group    Group  `json:"group"`
	Operator Member `json:"operator"`
}

type GroupAllowConfessTalk struct {
	EventBase
	Origin  bool  `json:"origin"`
	New     bool  `json:"new"`
	Current bool  `json:"current"`
	Group   Group `json:"group"`
	IsByBot bool  `json:"isByBot"`
}

type GroupAllowMemberInvite struct {
	EventBase
	Origin   bool   `json:"origin"`
	New      bool   `json:"new"`
	Current  bool   `json:"current"`
	Group    Group  `json:"group"`
	Operator Member `json:"operator"`
}

type MemberJoin struct {
	EventBase
	Member Member `json:"member"`
}

type MemberLeave struct {
	EventBase
	Member   Member `json:"member"`
	Operator Member `json:"operator,omitempty"`
}

type MemberCardChange struct {
	EventBase
	Origin   string `json:"origin"`
	New      string `json:"new"`
	Current  string `json:"current"`
	Member   Member `json:"member"`
	Operator Member `json:"operator"`
}

type MemberSpecialTitleChange struct {
	EventBase
	Origin  string `json:"origin"`
	New     string `json:"new"`
	Current string `json:"current"`
	Member  Member `json:"member"`
}

type MemberPermissionChange struct {
	EventBase
	Origin  GroupPermission `json:"origin"`
	New     GroupPermission `json:"new"`
	Current GroupPermission `json:"current"`
	Member  Member          `json:"member"`
}

type MemberMute struct {
	EventBase
	DurationSeconds int    `json:"durationSeconds"`
	Member          Member `json:"member"`
	Operator        Member `json:"operator"`
}

type MemberUnmute struct {
	EventBase
	Member   Member `json:"member"`
	Operator Member `json:"operator"`
}

type NewFriendRequest struct {
	EventBase
	EventId EventId `json:"eventId"`
	FromId  QQId    `json:"fromId"`
	GroupId GroupId `json:"groupId"`
	Nick    string  `json:"nick"`
	Message string  `json:"message"`
}

type MemberJoinRequest struct {
	EventBase
	EventId   EventId `json:"eventId"`
	FromId    QQId    `json:"fromId"`
	GroupId   GroupId `json:"groupId"`
	GroupName string  `json:"groupName"`
	Nick      string  `json:"nick"`
	Message   string  `json:"message"`
}

type BotInvitedJoinGroupRequest struct {
	EventBase
	EventId   EventId `json:"eventId"`
	FromId    QQId    `json:"fromId"`
	GroupId   GroupId `json:"groupId"`
	GroupName string  `json:"groupName"`
	Nick      string  `json:"nick"`
	Message   string  `json:"message"`
}

func (e *BotOnline) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *BotOffline) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *BotRelogin) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *GroupRecall) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *FriendRecall) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *BotGroupPermissionChange) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *BotMute) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *BotUnmute) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *BotJoinGroup) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *BotLeave) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *GroupNameChange) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *GroupEntranceAnnouncementChange) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *GroupMuteAll) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *GroupAllowAnonymousChat) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *GroupAllowConfessTalk) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *GroupAllowMemberInvite) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *MemberJoin) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *MemberLeave) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *MemberCardChange) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *MemberSpecialTitleChange) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *MemberPermissionChange) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *MemberMute) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *MemberUnmute) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *NewFriendRequest) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *MemberJoinRequest) String() string {
	s, _ := json.MarshalToString(e)
	return s
}
func (e *BotInvitedJoinGroupRequest) String() string {
	s, _ := json.MarshalToString(e)
	return s
}

func DeserializeEvent(rawjson []byte) (Event, error) {
	var eventType EventBase
	err := json.Unmarshal(rawjson, &eventType)
	if err != nil {
		return nil, err
	}

	var event Event
	switch eventType.Type {
	case BotOnlineEvent:
		var e BotOnline
		err = json.Unmarshal(rawjson, e)
		event = &e
	case BotOfflineEventActive, BotOfflineEventForce, BotOfflineEventDropped:
		var e BotOffline
		err = json.Unmarshal(rawjson, e)
		event = &e
	case BotReloginEvent:
		var e BotRelogin
		err = json.Unmarshal(rawjson, e)
		event = &e
	case GroupRecallEvent:
		var e GroupRecall
		err = json.Unmarshal(rawjson, e)
		event = &e
	case FriendRecallEvent:
		var e FriendRecall
		err = json.Unmarshal(rawjson, e)
		event = &e
	case BotGroupPermissionChangeEvent:
		var e BotGroupPermissionChange
		err = json.Unmarshal(rawjson, e)
		event = &e
	case BotMuteEvent:
		var e BotMute
		err = json.Unmarshal(rawjson, e)
		event = &e
	case BotUnmuteEvent:
		var e BotUnmute
		err = json.Unmarshal(rawjson, e)
		event = &e
	case BotJoinGroupEvent:
		var e BotJoinGroup
		err = json.Unmarshal(rawjson, e)
		event = &e
	case BotLeaveEventActive, BotLeaveEventKick:
		var e BotLeave
		err = json.Unmarshal(rawjson, e)
		event = &e
	case GroupNameChangeEvent:
		var e GroupNameChange
		err = json.Unmarshal(rawjson, e)
		event = &e
	case GroupEntranceAnnouncementChangeEvent:
		var e GroupEntranceAnnouncementChange
		err = json.Unmarshal(rawjson, e)
		event = &e
	case GroupMuteAllEvent:
		var e GroupMuteAll
		err = json.Unmarshal(rawjson, e)
		event = &e
	case GroupAllowAnonymousChatEvent:
		var e GroupAllowAnonymousChat
		err = json.Unmarshal(rawjson, e)
		event = &e
	case GroupAllowConfessTalkEvent:
		var e GroupAllowConfessTalk
		err = json.Unmarshal(rawjson, e)
		event = &e
	case GroupAllowMemberInviteEvent:
		var e GroupAllowMemberInvite
		err = json.Unmarshal(rawjson, e)
		event = &e
	case MemberJoinEvent:
		var e MemberJoin
		err = json.Unmarshal(rawjson, e)
		event = &e
	case MemberLeaveEventKick, MemberLeaveEventQuit:
		var e MemberLeave
		err = json.Unmarshal(rawjson, e)
		event = &e
	case MemberCardChangeEvent:
		var e MemberCardChange
		err = json.Unmarshal(rawjson, e)
		event = &e
	case MemberSpecialTitleChangeEvent:
		var e MemberSpecialTitleChange
		err = json.Unmarshal(rawjson, e)
		event = &e
	case MemberPermissionChangeEvent:
		var e MemberPermissionChange
		err = json.Unmarshal(rawjson, e)
		event = &e
	case MemberMuteEvent:
		var e MemberMute
		err = json.Unmarshal(rawjson, e)
		event = &e
	case MemberUnmuteEvent:
		var e MemberUnmute
		err = json.Unmarshal(rawjson, e)
		event = &e
	case NewFriendRequestEvent:
		var e NewFriendRequest
		err = json.Unmarshal(rawjson, e)
		event = &e
	case MemberJoinRequestEvent:
		var e MemberJoinRequest
		err = json.Unmarshal(rawjson, e)
		event = &e
	case BotInvitedJoinGroupRequestEvent:
		var e BotInvitedJoinGroupRequest
		err = json.Unmarshal(rawjson, e)
		event = &e
	default:
		err = errors.New("unknown event receive type")
	}
	if err != nil {
		return nil, err
	}
	return event, nil
}
