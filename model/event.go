package model

import (
	"github.com/wangnengjie/mirai-go/util/json"
)

type EventId int64

const (
	BotOnlineEvent                       MsgRecvType = "BotOnlineEvent"                       //Bot登录成功
	BotOfflineEventActive                MsgRecvType = "BotOfflineEventActive"                //Bot主动离线
	BotOfflineEventForce                 MsgRecvType = "BotOfflineEventForce"                 //Bot被挤下线
	BotOfflineEventDropped               MsgRecvType = "BotOfflineEventDropped"               //Bot被服务器断开或因网络问题而掉线
	BotReloginEvent                      MsgRecvType = "BotReloginEvent"                      //Bot主动重新登录
	GroupRecallEvent                     MsgRecvType = "GroupRecallEvent"                     //群消息撤回
	FriendRecallEvent                    MsgRecvType = "FriendRecallEvent"                    //好友消息撤回
	BotGroupPermissionChangeEvent        MsgRecvType = "BotGroupPermissionChangeEvent"        //Bot在群里的权限被改变. 操作人一定是群主
	BotMuteEvent                         MsgRecvType = "BotMuteEvent"                         //Bot被禁言
	BotUnmuteEvent                       MsgRecvType = "BotUnmuteEvent"                       //Bot被取消禁言
	BotJoinGroupEvent                    MsgRecvType = "BotJoinGroupEvent"                    //Bot加入了一个新群
	BotLeaveEventActive                  MsgRecvType = "BotLeaveEventActive"                  //Bot主动退出一个群
	BotLeaveEventKick                    MsgRecvType = "BotLeaveEventKick"                    //Bot被踢出一个群
	GroupNameChangeEvent                 MsgRecvType = "GroupNameChangeEvent"                 //某个群名改变
	GroupEntranceAnnouncementChangeEvent MsgRecvType = "GroupEntranceAnnouncementChangeEvent" //某群入群公告改变
	GroupMuteAllEvent                    MsgRecvType = "GroupMuteAllEvent"                    //全员禁言
	GroupAllowAnonymousChatEvent         MsgRecvType = "GroupAllowAnonymousChatEvent"         //匿名聊天
	GroupAllowConfessTalkEvent           MsgRecvType = "GroupAllowConfessTalkEvent"           //坦白说
	GroupAllowMemberInviteEvent          MsgRecvType = "GroupAllowMemberInviteEvent"          //允许群员邀请好友加群
	MemberJoinEvent                      MsgRecvType = "MemberJoinEvent"                      //新人入群的事件
	MemberLeaveEventKick                 MsgRecvType = "MemberLeaveEventKick"                 //成员被踢出群（该成员不是Bot）
	MemberLeaveEventQuit                 MsgRecvType = "MemberLeaveEventQuit"                 //成员主动离群（该成员不是Bot）
	MemberCardChangeEvent                MsgRecvType = "MemberCardChangeEvent"                //群名片改动
	MemberSpecialTitleChangeEvent        MsgRecvType = "MemberSpecialTitleChangeEvent"        //群头衔改动（只有群主有操作限权）
	MemberPermissionChangeEvent          MsgRecvType = "MemberPermissionChangeEvent"          //成员权限改变的事件（该成员不可能是Bot，见BotGroupPermissionChangeEvent）
	MemberMuteEvent                      MsgRecvType = "MemberMuteEvent"                      //群成员被禁言事件（该成员不可能是Bot，见BotMuteEvent）
	MemberUnmuteEvent                    MsgRecvType = "MemberUnmuteEvent"                    //群成员被取消禁言事件（该成员不可能是Bot，见BotUnmuteEvent）
	NewFriendRequestEvent                MsgRecvType = "NewFriendRequestEvent"                //添加好友申请
	MemberJoinRequestEvent               MsgRecvType = "MemberJoinRequestEvent"               //用户入群申请（Bot需要有管理员权限）
	BotInvitedJoinGroupRequestEvent      MsgRecvType = "BotInvitedJoinGroupRequestEvent"      //Bot被邀请入群申请
)

type (
	BotOnline struct {
		MsgRecvBase
		QQ QQId `json:"qq"`
	}

	BotOffline struct {
		MsgRecvBase
		QQ QQId `json:"qq"`
	}

	BotRelogin struct {
		MsgRecvBase
		QQ QQId `json:"qq"`
	}

	GroupRecall struct {
		MsgRecvBase
		AuthorId  QQId   `json:"authorId"`
		MessageId MsgId  `json:"messageId"`
		Time      int    `json:"time"`
		Group     Group  `json:"group"`
		Operator  Member `json:"operator"`
	}

	FriendRecall struct {
		MsgRecvBase
		AuthorId  QQId  `json:"authorId"`
		MessageId MsgId `json:"messageId"`
		Time      int   `json:"time"`
		Operator  QQId  `json:"operator"`
	}

	BotGroupPermissionChange struct {
		MsgRecvBase
		Origin  GroupPermission `json:"origin"`
		New     GroupPermission `json:"new"`
		Current GroupPermission `json:"current"`
		Group   Group           `json:"group"`
	}

	BotMute struct {
		MsgRecvBase
		DurationSeconds int    `json:"durationSeconds"`
		Operator        Member `json:"operator"`
	}

	BotUnmute struct {
		MsgRecvBase
		Operator Member `json:"operator"`
	}

	BotJoinGroup struct {
		MsgRecvBase
		Group Group `json:"group"`
	}

	BotLeave struct {
		MsgRecvBase
		Group Group `json:"group"`
	}

	GroupNameChange struct {
		MsgRecvBase
		Origin   string `json:"origin"`
		New      string `json:"new"`
		Current  string `json:"current"`
		Group    Group  `json:"group"`
		Operator Member `json:"operator"`
	}

	GroupEntranceAnnouncementChange struct {
		MsgRecvBase
		Origin   string `json:"origin"`
		New      string `json:"new"`
		Current  string `json:"current"`
		Group    Group  `json:"group"`
		Operator Member `json:"operator"`
	}

	GroupMuteAll struct {
		MsgRecvBase
		Origin   bool   `json:"origin"`
		New      bool   `json:"new"`
		Current  bool   `json:"current"`
		Group    Group  `json:"group"`
		Operator Member `json:"operator"`
	}

	GroupAllowAnonymousChat struct {
		MsgRecvBase
		Origin   bool   `json:"origin"`
		New      bool   `json:"new"`
		Current  bool   `json:"current"`
		Group    Group  `json:"group"`
		Operator Member `json:"operator"`
	}

	GroupAllowConfessTalk struct {
		MsgRecvBase
		Origin  bool  `json:"origin"`
		New     bool  `json:"new"`
		Current bool  `json:"current"`
		Group   Group `json:"group"`
		IsByBot bool  `json:"isByBot"`
	}

	GroupAllowMemberInvite struct {
		MsgRecvBase
		Origin   bool   `json:"origin"`
		New      bool   `json:"new"`
		Current  bool   `json:"current"`
		Group    Group  `json:"group"`
		Operator Member `json:"operator"`
	}

	MemberJoin struct {
		MsgRecvBase
		Member Member `json:"member"`
	}

	MemberLeave struct {
		MsgRecvBase
		Member   Member `json:"member"`
		Operator Member `json:"operator,omitempty"`
	}

	MemberCardChange struct {
		MsgRecvBase
		Origin   string `json:"origin"`
		New      string `json:"new"`
		Current  string `json:"current"`
		Member   Member `json:"member"`
		Operator Member `json:"operator"`
	}

	MemberSpecialTitleChange struct {
		MsgRecvBase
		Origin  string `json:"origin"`
		New     string `json:"new"`
		Current string `json:"current"`
		Member  Member `json:"member"`
	}

	MemberPermissionChange struct {
		MsgRecvBase
		Origin  GroupPermission `json:"origin"`
		New     GroupPermission `json:"new"`
		Current GroupPermission `json:"current"`
		Member  Member          `json:"member"`
	}

	MemberMute struct {
		MsgRecvBase
		DurationSeconds int    `json:"durationSeconds"`
		Member          Member `json:"member"`
		Operator        Member `json:"operator"`
	}

	MemberUnmute struct {
		MsgRecvBase
		Member   Member `json:"member"`
		Operator Member `json:"operator"`
	}

	NewFriendRequest struct {
		MsgRecvBase
		EventId EventId `json:"eventId"`
		FromId  QQId    `json:"fromId"`
		GroupId GroupId `json:"groupId"`
		Nick    string  `json:"nick"`
		Message string  `json:"message"`
	}

	MemberJoinRequest struct {
		MsgRecvBase
		EventId   EventId `json:"eventId"`
		FromId    QQId    `json:"fromId"`
		GroupId   GroupId `json:"groupId"`
		GroupName string  `json:"groupName"`
		Nick      string  `json:"nick"`
		Message   string  `json:"message"`
	}

	BotInvitedJoinGroupRequest struct {
		MsgRecvBase
		EventId   EventId `json:"eventId"`
		FromId    QQId    `json:"fromId"`
		GroupId   GroupId `json:"groupId"`
		GroupName string  `json:"groupName"`
		Nick      string  `json:"nick"`
		Message   string  `json:"message"`
	}
)

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
