package model

type NoticeViewModel struct {
	Notice
	SendCount int `:"SendCount"` // 发送条数
	Count     int `:"Count"`     // 总条数
}
