package model

import "github.com/wangnengjie/mirai-go/util/json"

type GroupPermission string

const (
	OWNER         GroupPermission = "OWNER"
	ADMINISTRATOR GroupPermission = "ADMINISTRATOR"
	MEMBER        GroupPermission = "MEMBER"
)

type (
	User struct {
		Id       QQId   `json:"id"`
		NickName string `json:"nickName"`
		Remark   string `json:"remark"`
	}

	Group struct {
		Id         GroupId         `json:"id"`
		Name       string          `json:"name"`
		Permission GroupPermission `json:"permission"`
	}

	Member struct {
		Id         QQId            `json:"id"`
		MemberName string          `json:"memberName"`
		Permission GroupPermission `json:"permission"`
		Group      Group           `json:"group"`
	}

	MemberInfo struct {
		Name         string `json:"name,omitempty"`
		SpecialTitle string `json:"specialTitle,omitempty"`
	}

	GroupConfig struct {
		Name              string `json:"name,omitempty"`
		Announcement      string `json:"announcement,omitempty"`
		ConfessTalk       *bool  `json:"confessTalk,omitempty"`
		AllowMemberInvite *bool  `json:"allowMemberInvite,omitempty"`
		AutoApprove       *bool  `json:"autoApprove,omitempty"`
		AnonymousChat     *bool  `json:"anonymousChat,omitempty"`
	}
)

func (m *User) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *Group) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *Member) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *MemberInfo) String() string {
	s, _ := json.MarshalToString(m)
	return s
}

func (m *GroupConfig) String() string {
	s, _ := json.MarshalToString(m)
	return s
}
