package model
type User struct {
	Model
	AccountName  string
	PassWord  string
	RealName  string
	MobilePhone  string
	Email  string
	Description  string
	IsAble int64
	IfChangePwd int64
}
