package model

// type User struct {
// 	ID          string `:"ID"`          // 主键
// 	AccountName string `:"AccountName"` // 帐户名
// 	PassWord    string `:"PassWord"`    // 帐户密码（32位MD5加密）
// 	RealName    string `:"RealName"`    // 别名
// 	MobilePhone string `:"MobilePhone"` // 联系方式
// 	Email       string `:"Email"`       // 邮箱
// 	Description string `:"Description"` // 简介
// 	CreateBy    string `:"CreateBy"`    // 创建人
// 	CreateTime  string `:"CreateTime"`  // 创建时间
// 	UpdateBy    string `:"UpdateBy"`    // 修改人
// 	UpdateTime  string `:"UpdateTime"`  // 修改时间
// 	IsAble      int    `:"IsAble"`      // 是否有效
// 	IfChangePwd int    `:"IfChangePwd"` // 是否修改密码
// }
type TokenUser struct {
	ID          string `:"ID"`          // 主键
	AccountName string `:"AccountName"` // 帐户名
	PassWord    string `:"PassWord"`    // 帐户密码（32位MD5加密）
	RealName    string `:"RealName"`    // 别名
	MobilePhone string `:"MobilePhone"` // 联系方式
	Email       string `:"Email"`       // 邮箱
	Description string `:"Description"` // 简介
	CreateBy    string `:"CreateBy"`    // 创建人
	CreateTime  string `:"CreateTime"`  // 创建时间
	UpdateBy    string `:"UpdateBy"`    // 修改人
	UpdateTime  string `:"UpdateTime"`  // 修改时间
	IsAble      int    `:"IsAble"`      // 是否有效
	IfChangePwd int    `:"IfChangePwd"` // 是否修改密码
	Token       string
}

type UserView struct {
	ID             string `:"ID"`          // 主键
	AccountName    string `:"AccountName"` // 帐户名
	PassWord       string `:"PassWord"`    // 帐户密码（32位MD5加密）
	RealName       string `:"RealName"`    // 别名
	MobilePhone    string `:"MobilePhone"` // 联系方式
	Email          string `:"Email"`       // 邮箱
	Description    string `:"Description"` // 简介
	CreateBy       string `:"CreateBy"`    // 创建人
	CreateTime     string `:"CreateTime"`  // 创建时间
	UpdateBy       string `:"UpdateBy"`    // 修改人
	UpdateTime     string `:"UpdateTime"`  // 修改时间
	IsAble         int    `:"IsAble"`      // 是否有效
	IfChangePwd    int    `:"IfChangePwd"` // 是否修改密码
	DepartmentName string `:"DepartmentName"`
	RoleName       string `:"DepartmRoleNameentName"`
}

