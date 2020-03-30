package model

type NoticeViewModel struct {
	Notice
	SendCount int `:"SendCount"` // 发送条数
	Count     int `:"Count"`     // 总条数
}
type NoticeBillModel struct {
	Main Notice
	Item []NoticeUserView
}

// type NoticeUserView struct {
// 	// NoticeUser
// 	ID          string `:"ID"`          // 主键
// 	NoticeId    string `:"NoticeId"`    // 通知单ID
// 	UserId      string `:"UserId"`      // 用户ID
// 	Notes       string `:"Notes"`       // 备注
// 	CreateBy    string `:"CreateBy"`    // 创建人
// 	CreateTime  string `:"CreateTime"`  // 创建时间
// 	UpdateBy    string `:"UpdateBy"`    // 修改人
// 	UpdateTime  string `:"UpdateTime"`  // 修改时间
// 	SendFlag    int    `:"SendFlag"`    // 发送标记
// 	SendTime    string `:"SendTime"`    // 发送时间
// 	AccountName string `:"AccountName"` // 帐户名
// 	RealName    string `:"RealName"`    // 别名
// }
