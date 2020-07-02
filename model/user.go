package model

type GroupPermission string

const (
	OWNER         GroupPermission = "OWNER"
	ADMINISTRATOR GroupPermission = "ADMINISTRATOR"
	MEMBER        GroupPermission = "MEMBER"
)

type User struct {
	Id       QQId   `json:"id"`
	NickName string `json:"nickName"`
	Remark   string `json:"remark"`
}

type Group struct {
	Id         GroupId         `json:"id"`
	Name       string          `json:"name"`
	Permission GroupPermission `json:"permission"`
}

type Member struct {
	Id         QQId            `json:"id"`
	MemberName string          `json:"memberName"`
	Permission GroupPermission `json:"permission"`
	Group      Group           `json:"group"`
}
