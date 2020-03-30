package model

import (
	"time"
)

type Model struct {
	ID         int64
	CreateBy   string
	UpdateBy   string
	CreateTime time.Time
	UpdateTime time.Time
}
type FilterModel struct {
	Column   string `json:"column"`   //字段名
	Action   string `json:"action"`   //操作符 > < =
	Logic    string `json:"logic"`    //关系  and or
	Value    string `json:"value"`    //值
	DataType string `json:"dataType"` //数据类型
}
type TransferModel struct {
	Key      string `json:"key"`
	Label    string `json:"label"`
	Title    string `json:"title"`
	Disabled bool   `json:"disabled"`
}
type TreeModel struct {
	ID       string      `json:"id"`
	Label    string      `json:"label"`
	RealName string      `json:"RealName"`
	Children []TreeModel `json:"children"`
}
type TreeNodeModel struct {
	ID       string          `json:"id"`
	Value    string          `json:"value"`
	Key      string          `json:"key"`
	Label    string          `json:"label"`
	RealName string          `json:"RealName"`
	Children []TreeNodeModel `json:"children"`
}
type CascaderListModel struct {
	Value string `json:"value"` //值
	Label string `json:"label"`
}

type WsInfoModel struct {
	Data string `json:"Data"`
	Type string `json:"Type"`
	ID   string `json:"ID"`
}
type MsgInfo struct {
	ID       string `:"ID"`       // 主键
	No       string `:"No"`       // 单号
	MsgType  string `:"MsgType"`  // 消息类型
	TypeId   int    `:"TypeId"`   // 通知类型
	Title    string `:"Title"`    // 标题
	Content  string `:"Content"`  // 内容
	OpenFlag int    `:"OpenFlag"` // 打开标记
	BillTime string `:"BillTime"` // 消息生成时间
	MsgTime  string `:"MsgTime"`  // 消息发送时间

}

// model

type NoticeUser struct {
	ID         string `orm:"ID" json:"ID"`                 // 主键
	NoticeId   string `orm:"NoticeId" json:"NoticeId"`     // 通知单ID
	UserId     string `orm:"UserId" json:"UserId"`         // 用户ID
	Notes      string `orm:"Notes" json:"Notes"`           // 备注
	CreateBy   string `orm:"CreateBy" json:"CreateBy"`     // 创建人
	CreateTime string `orm:"CreateTime" json:"CreateTime"` // 创建时间
	UpdateBy   string `orm:"UpdateBy" json:"UpdateBy"`     // 修改人
	UpdateTime string `orm:"UpdateTime" json:"UpdateTime"` // 修改时间
	SendFlag   int    `orm:"SendFlag" json:"SendFlag"`     // 发送标记
	SendTime   string `orm:"SendTime" json:"SendTime"`     // 发送时间
}

func (*NoticeUser) TableName() string {
	return "NoticeUser"
}

type User struct {
	ID          string `orm:"ID" json:"ID"`                   // 主键
	AccountName string `orm:"AccountName" json:"AccountName"` // 帐户名
	PassWord    string `orm:"PassWord" json:"PassWord"`       // 帐户密码（32位MD5加密）
	RealName    string `orm:"RealName" json:"RealName"`       // 别名
	MobilePhone string `orm:"MobilePhone" json:"MobilePhone"` // 联系方式
	Email       string `orm:"Email" json:"Email"`             // 邮箱
	Description string `orm:"Description" json:"Description"` // 简介
	CreateBy    string `orm:"CreateBy" json:"CreateBy"`       // 创建人
	CreateTime  string `orm:"CreateTime" json:"CreateTime"`   // 创建时间
	UpdateBy    string `orm:"UpdateBy" json:"UpdateBy"`       // 修改人
	UpdateTime  string `orm:"UpdateTime" json:"UpdateTime"`   // 修改时间
	IsAble      int    `orm:"IsAble" json:"IsAble"`           // 是否有效
	IfChangePwd int    `orm:"IfChangePwd" json:"IfChangePwd"` // 是否修改密码
}

func (*User) TableName() string {
	return "User"
}

type Button struct {
	ID          string `orm:"ID" json:"ID"`                   // 主键
	Name        string `orm:"Name" json:"Name"`               // 按钮名
	Code        string `orm:"Code" json:"Code"`               // 按钮代码
	Icon        string `orm:"Icon" json:"Icon"`               // 按钮图标
	Sort        int    `orm:"Sort" json:"Sort"`               // 排序
	Description string `orm:"Description" json:"Description"` // 简介
	CreateBy    string `orm:"CreateBy" json:"CreateBy"`       // 创建人
	CreateTime  string `orm:"CreateTime" json:"CreateTime"`   // 创建时间
	UpdateBy    string `orm:"UpdateBy" json:"UpdateBy"`       // 修改人
	UpdateTime  string `orm:"UpdateTime" json:"UpdateTime"`   // 修改时间
	IsAble      int    `orm:"IsAble" json:"IsAble"`           // 是否有效
}

func (*Button) TableName() string {
	return "Button"
}

type DeptView struct {
	ID          string `orm:"ID" json:"ID"`                   // 部门ID
	Name        string `orm:"Name" json:"Name"`               // 部门名
	ParentId    string `orm:"ParentId" json:"ParentId"`       // 父级部门名
	Sort        int    `orm:"Sort" json:"Sort"`               // 排序
	Description string `orm:"Description" json:"Description"` // 简介
	CreateBy    string `orm:"CreateBy" json:"CreateBy"`       // 创建人
	CreateTime  string `orm:"CreateTime" json:"CreateTime"`   // 创建时间
	UpdateBy    string `orm:"UpdateBy" json:"UpdateBy"`       // 修改人
	UpdateTime  string `orm:"UpdateTime" json:"UpdateTime"`   // 修改时间
	IsAble      int    `orm:"IsAble" json:"IsAble"`           // 是否有效
	ParentName  string `orm:"ParentName" json:"ParentName"`
}

func (*DeptView) TableName() string {
	return "DeptView"
}

type MenuView struct {
	ID          string `orm:"ID" json:"ID"`                   // 主键
	Name        string `orm:"Name" json:"Name"`               // 菜单名
	ParentId    string `orm:"ParentId" json:"ParentId"`       // 父级菜单名
	Code        string `orm:"Code" json:"Code"`               // 菜单代码
	LinkAddress string `orm:"LinkAddress" json:"LinkAddress"` // 菜单地址
	Icon        string `orm:"Icon" json:"Icon"`               // 菜单图标
	Sort        int    `orm:"Sort" json:"Sort"`               // 排序
	Description string `orm:"Description" json:"Description"` // 简介
	CreateBy    string `orm:"CreateBy" json:"CreateBy"`       // 创建人
	CreateTime  string `orm:"CreateTime" json:"CreateTime"`   // 创建时间
	UpdateBy    string `orm:"UpdateBy" json:"UpdateBy"`       // 修改人
	UpdateTime  string `orm:"UpdateTime" json:"UpdateTime"`   // 修改时间
	IsAble      int    `orm:"IsAble" json:"IsAble"`           // 是否有效
	ParentName  string `orm:"ParentName" json:"ParentName"`
}

func (*MenuView) TableName() string {
	return "MenuView"
}

type Role struct {
	ID          string `orm:"ID" json:"ID"`                   // 主键
	Name        string `orm:"Name" json:"Name"`               // 权限名
	Description string `orm:"Description" json:"Description"` // 简介
	CreateBy    string `orm:"CreateBy" json:"CreateBy"`       // 创建人
	CreateTime  string `orm:"CreateTime" json:"CreateTime"`   // 创建时间
	UpdateBy    string `orm:"UpdateBy" json:"UpdateBy"`       // 修改人
	UpdateTime  string `orm:"UpdateTime" json:"UpdateTime"`   // 修改时间
	IsAble      int    `orm:"IsAble" json:"IsAble"`           // 是否有效
}

func (*Role) TableName() string {
	return "Role"
}

type UserRole struct {
	ID         string `orm:"ID" json:"ID"`                 // 部门ID
	UserId     string `orm:"UserId" json:"UserId"`         // 用户ID
	RoleId     string `orm:"RoleId" json:"RoleId"`         // 权限ID
	CreateBy   string `orm:"CreateBy" json:"CreateBy"`     // 创建人
	CreateTime string `orm:"CreateTime" json:"CreateTime"` // 创建时间
	UpdateBy   string `orm:"UpdateBy" json:"UpdateBy"`     // 修改人
	UpdateTime string `orm:"UpdateTime" json:"UpdateTime"` // 修改时间
}

func (*UserRole) TableName() string {
	return "UserRole"
}

type UserToken struct {
	Token        string `orm:"Token" json:"Token"`               // 主键
	UserId       string `orm:"UserId" json:"UserId"`             // 用户ID
	Port         string `orm:"Port" json:"Port"`                 // 端口
	Address      string `orm:"Address" json:"Address"`           // 地址
	CreateTime   string `orm:"CreateTime" json:"CreateTime"`     // 创建时间
	UpdateTime   string `orm:"UpdateTime" json:"UpdateTime"`     // 修改时间
	IsLoginOut   int    `orm:"IsLoginOut" json:"IsLoginOut"`     // 离线标记
	LoginOutTime string `orm:"LoginOutTime" json:"LoginOutTime"` // 离线时间
}

func (*UserToken) TableName() string {
	return "UserToken"
}

type NoticeUserView struct {
	ID          string `orm:"ID" json:"ID"`                   // 主键
	NoticeId    string `orm:"NoticeId" json:"NoticeId"`       // 通知单ID
	UserId      string `orm:"UserId" json:"UserId"`           // 用户ID
	Notes       string `orm:"Notes" json:"Notes"`             // 备注
	CreateBy    string `orm:"CreateBy" json:"CreateBy"`       // 创建人
	CreateTime  string `orm:"CreateTime" json:"CreateTime"`   // 创建时间
	UpdateBy    string `orm:"UpdateBy" json:"UpdateBy"`       // 修改人
	UpdateTime  string `orm:"UpdateTime" json:"UpdateTime"`   // 修改时间
	SendFlag    int    `orm:"SendFlag" json:"SendFlag"`       // 发送标记
	SendTime    string `orm:"SendTime" json:"SendTime"`       // 发送时间
	AccountName string `orm:"AccountName" json:"AccountName"` // 帐户名
	RealName    string `orm:"RealName" json:"RealName"`       // 别名
}

func (*NoticeUserView) TableName() string {
	return "NoticeUserView"
}

type UserDepartment struct {
	ID           string `orm:"ID" json:"ID"`                     // 部门ID
	UserId       string `orm:"UserId" json:"UserId"`             // 用户ID
	DepartmentId string `orm:"DepartmentId" json:"DepartmentId"` // 部门ID
	CreateBy     string `orm:"CreateBy" json:"CreateBy"`         // 创建人
	CreateTime   string `orm:"CreateTime" json:"CreateTime"`     // 创建时间
	UpdateBy     string `orm:"UpdateBy" json:"UpdateBy"`         // 修改人
	UpdateTime   string `orm:"UpdateTime" json:"UpdateTime"`     // 修改时间
}

func (*UserDepartment) TableName() string {
	return "UserDepartment"
}

type BillNo struct {
	ID            string `orm:"ID" json:"ID"`                       // ID
	Code          string `orm:"Code" json:"Code"`                   // 单位据代码
	CodeName      string `orm:"CodeName" json:"CodeName"`           // 单据号名称
	CurrentId     int    `orm:"CurrentId" json:"CurrentId"`         // 当前流水号
	CurrentBillNo string `orm:"CurrentBillNo" json:"CurrentBillNo"` // 当前单据号
	CurrentTime   string `orm:"CurrentTime" json:"CurrentTime"`     // 当前单据生成时间
	MaskInfo      string `orm:"MaskInfo" json:"MaskInfo"`           // 掩码
	EndLength     int    `orm:"EndLength" json:"EndLength"`         // 流水号长度
	CreateBy      string `orm:"CreateBy" json:"CreateBy"`           // 创建人
	CreateTime    string `orm:"CreateTime" json:"CreateTime"`       // 创建时间
	UpdateBy      string `orm:"UpdateBy" json:"UpdateBy"`           // 修改人
	UpdateTime    string `orm:"UpdateTime" json:"UpdateTime"`       // 修改时间
}

func (*BillNo) TableName() string {
	return "BillNo"
}

type Notice struct {
	ID            string `orm:"ID" json:"ID"`                       // 主键
	No            string `orm:"No" json:"No"`                       // 通知单号
	TypeId        int    `orm:"TypeId" json:"TypeId"`               // 通知类型
	Title         string `orm:"Title" json:"Title"`                 // 通知标题
	Content       string `orm:"Content" json:"Content"`             // 通知内容
	Notes         string `orm:"Notes" json:"Notes"`                 // 备注
	CreateBy      string `orm:"CreateBy" json:"CreateBy"`           // 创建人
	CreateTime    string `orm:"CreateTime" json:"CreateTime"`       // 创建时间
	UpdateBy      string `orm:"UpdateBy" json:"UpdateBy"`           // 修改人
	UpdateTime    string `orm:"UpdateTime" json:"UpdateTime"`       // 修改时间
	Status        int    `orm:"Status" json:"Status"`               // 通知状态
	NoticeTime    string `orm:"NoticeTime" json:"NoticeTime"`       // 通知时间
	SendBeginTime string `orm:"SendBeginTime" json:"SendBeginTime"` // 通知发送开始时间
	SendEndTime   string `orm:"SendEndTime" json:"SendEndTime"`     // 通知发送结束时间
}

func (*Notice) TableName() string {
	return "Notice"
}

type NoticeView struct {
	ID            string `orm:"ID" json:"ID"`                       // 主键
	No            string `orm:"No" json:"No"`                       // 通知单号
	TypeId        int    `orm:"TypeId" json:"TypeId"`               // 通知类型
	Title         string `orm:"Title" json:"Title"`                 // 通知标题
	Content       string `orm:"Content" json:"Content"`             // 通知内容
	Notes         string `orm:"Notes" json:"Notes"`                 // 备注
	CreateBy      string `orm:"CreateBy" json:"CreateBy"`           // 创建人
	CreateTime    string `orm:"CreateTime" json:"CreateTime"`       // 创建时间
	UpdateBy      string `orm:"UpdateBy" json:"UpdateBy"`           // 修改人
	UpdateTime    string `orm:"UpdateTime" json:"UpdateTime"`       // 修改时间
	Status        int    `orm:"Status" json:"Status"`               // 通知状态
	NoticeTime    string `orm:"NoticeTime" json:"NoticeTime"`       // 通知时间
	SendBeginTime string `orm:"SendBeginTime" json:"SendBeginTime"` // 通知发送开始时间
	SendEndTime   string `orm:"SendEndTime" json:"SendEndTime"`     // 通知发送结束时间
	SendCount     int    `orm:"SendCount" json:"SendCount"`
	Count         int    `orm:"Count" json:"Count"`
}

func (*NoticeView) TableName() string {
	return "NoticeView"
}

type Menu struct {
	ID          string `orm:"ID" json:"ID"`                   // 主键
	Name        string `orm:"Name" json:"Name"`               // 菜单名
	ParentId    string `orm:"ParentId" json:"ParentId"`       // 父级菜单名
	Code        string `orm:"Code" json:"Code"`               // 菜单代码
	LinkAddress string `orm:"LinkAddress" json:"LinkAddress"` // 菜单地址
	Icon        string `orm:"Icon" json:"Icon"`               // 菜单图标
	Sort        int    `orm:"Sort" json:"Sort"`               // 排序
	Description string `orm:"Description" json:"Description"` // 简介
	CreateBy    string `orm:"CreateBy" json:"CreateBy"`       // 创建人
	CreateTime  string `orm:"CreateTime" json:"CreateTime"`   // 创建时间
	UpdateBy    string `orm:"UpdateBy" json:"UpdateBy"`       // 修改人
	UpdateTime  string `orm:"UpdateTime" json:"UpdateTime"`   // 修改时间
	IsAble      int    `orm:"IsAble" json:"IsAble"`           // 是否有效
}

func (*Menu) TableName() string {
	return "Menu"
}

type Parameter struct {
	ID          string `orm:"ID" json:"ID"`                   // 主键
	Name        string `orm:"Name" json:"Name"`               // 参数名称
	Code        string `orm:"Code" json:"Code"`               // 参数代码
	ParentId    string `orm:"ParentId" json:"ParentId"`       // 上级参数
	Sort        int    `orm:"Sort" json:"Sort"`               // 排序
	Description string `orm:"Description" json:"Description"` // 简介
	CreateBy    string `orm:"CreateBy" json:"CreateBy"`       // 创建人
	CreateTime  string `orm:"CreateTime" json:"CreateTime"`   // 创建时间
	UpdateBy    string `orm:"UpdateBy" json:"UpdateBy"`       // 修改人
	UpdateTime  string `orm:"UpdateTime" json:"UpdateTime"`   // 修改时间
	IsAble      int    `orm:"IsAble" json:"IsAble"`           // 是否有效
}

func (*Parameter) TableName() string {
	return "Parameter"
}

type FlowLog struct {
	ID          string `orm:"ID" json:"ID"`                   // 主键
	FlowId      string `orm:"FlowId" json:"FlowId"`           // 流程ID
	BillId      string `orm:"BillId" json:"BillId"`           // 单据ID
	IsApproval  int    `orm:"IsApproval" json:"IsApproval"`   // 是否审批通过
	Description string `orm:"Description" json:"Description"` // 简介
	Notes       string `orm:"Notes" json:"Notes"`             // 备注
	CreateBy    string `orm:"CreateBy" json:"CreateBy"`       // 创建人
	CreateTime  string `orm:"CreateTime" json:"CreateTime"`   // 创建时间
	UpdateBy    string `orm:"UpdateBy" json:"UpdateBy"`       // 修改人
	UpdateTime  string `orm:"UpdateTime" json:"UpdateTime"`   // 修改时间
	Status      int    `orm:"Status" json:"Status"`           // 状态
}

func (*FlowLog) TableName() string {
	return "FlowLog"
}

type FlowStepUser struct {
	ID         string `orm:"ID" json:"ID"`                 // 主键
	FlowId     string `orm:"FlowId" json:"FlowId"`         // 流程ID
	StepId     string `orm:"StepId" json:"StepId"`         // 步骤ID
	UserId     string `orm:"UserId" json:"UserId"`         // 用户ID
	Notes      string `orm:"Notes" json:"Notes"`           // 备注
	CreateBy   string `orm:"CreateBy" json:"CreateBy"`     // 创建人
	CreateTime string `orm:"CreateTime" json:"CreateTime"` // 创建时间
	UpdateBy   string `orm:"UpdateBy" json:"UpdateBy"`     // 修改人
	UpdateTime string `orm:"UpdateTime" json:"UpdateTime"` // 修改时间
}

func (*FlowStepUser) TableName() string {
	return "FlowStepUser"
}

type FlowStepUserLog struct {
	ID         string `orm:"ID" json:"ID"`                 // 主键
	FlowLogId  string `orm:"FlowLogId" json:"FlowLogId"`   // 流程日志ID
	StepLogId  string `orm:"StepLogId" json:"StepLogId"`   // 步骤日志ID
	FlowId     string `orm:"FlowId" json:"FlowId"`         // 流程ID
	StepId     string `orm:"StepId" json:"StepId"`         // 步骤ID
	UserId     string `orm:"UserId" json:"UserId"`         // 用户ID
	Notes      string `orm:"Notes" json:"Notes"`           // 备注
	CreateBy   string `orm:"CreateBy" json:"CreateBy"`     // 创建人
	CreateTime string `orm:"CreateTime" json:"CreateTime"` // 创建时间
	UpdateBy   string `orm:"UpdateBy" json:"UpdateBy"`     // 修改人
	UpdateTime string `orm:"UpdateTime" json:"UpdateTime"` // 修改时间
	IsApproval int    `orm:"IsApproval" json:"IsApproval"` // 是否审批通过
}

func (*FlowStepUserLog) TableName() string {
	return "FlowStepUserLog"
}

type FlowStepUserView struct {
	ID          string `orm:"ID" json:"ID"`                 // 主键
	FlowId      string `orm:"FlowId" json:"FlowId"`         // 流程ID
	StepId      string `orm:"StepId" json:"StepId"`         // 步骤ID
	UserId      string `orm:"UserId" json:"UserId"`         // 用户ID
	Notes       string `orm:"Notes" json:"Notes"`           // 备注
	CreateBy    string `orm:"CreateBy" json:"CreateBy"`     // 创建人
	CreateTime  string `orm:"CreateTime" json:"CreateTime"` // 创建时间
	UpdateBy    string `orm:"UpdateBy" json:"UpdateBy"`     // 修改人
	UpdateTime  string `orm:"UpdateTime" json:"UpdateTime"` // 修改时间
	AccountName string `orm:"AccountName" json:"AccountName"`
	RealName    string `orm:"RealName" json:"RealName"`
}

func (*FlowStepUserView) TableName() string {
	return "FlowStepUserView"
}

type FlowView struct {
	ID          string `orm:"ID" json:"ID"`                   // 主键
	No          string `orm:"No" json:"No"`                   // 单据号
	MenuId      string `orm:"MenuId" json:"MenuId"`           // 菜单ID
	Description string `orm:"Description" json:"Description"` // 简介
	Notes       string `orm:"Notes" json:"Notes"`             // 备注
	CreateBy    string `orm:"CreateBy" json:"CreateBy"`       // 创建人
	CreateTime  string `orm:"CreateTime" json:"CreateTime"`   // 创建时间
	UpdateBy    string `orm:"UpdateBy" json:"UpdateBy"`       // 修改人
	UpdateTime  string `orm:"UpdateTime" json:"UpdateTime"`   // 修改时间
	Status      int    `orm:"Status" json:"Status"`           // 状态
	MenuCode    string `orm:"MenuCode" json:"MenuCode"`       // 菜单代码
	MenuName    string `orm:"MenuName" json:"MenuName"`       // 菜单名
	StepNum     int    `orm:"StepNum" json:"StepNum"`
}

func (*FlowView) TableName() string {
	return "FlowView"
}

type Icons struct {
	ID         string `orm:"ID" json:"ID"`                 // ID
	Code       string `orm:"Code" json:"Code"`             // 图标代码
	CodeName   string `orm:"CodeName" json:"CodeName"`     // 图标名称
	CssInfo    string `orm:"CssInfo" json:"CssInfo"`       // 图标样式
	CreateBy   string `orm:"CreateBy" json:"CreateBy"`     // 创建人
	CreateTime string `orm:"CreateTime" json:"CreateTime"` // 创建时间
	UpdateBy   string `orm:"UpdateBy" json:"UpdateBy"`     // 修改人
	UpdateTime string `orm:"UpdateTime" json:"UpdateTime"` // 修改时间
}

func (*Icons) TableName() string {
	return "Icons"
}

type RoleMenu struct {
	ID         string `orm:"ID" json:"ID"`                 // ID
	RoleId     string `orm:"RoleId" json:"RoleId"`         // 权限ID
	MenuId     string `orm:"MenuId" json:"MenuId"`         // 菜单ID
	CreateBy   string `orm:"CreateBy" json:"CreateBy"`     // 创建人
	CreateTime string `orm:"CreateTime" json:"CreateTime"` // 创建时间
	UpdateBy   string `orm:"UpdateBy" json:"UpdateBy"`     // 修改人
	UpdateTime string `orm:"UpdateTime" json:"UpdateTime"` // 修改时间
}

func (*RoleMenu) TableName() string {
	return "RoleMenu"
}

type RoleMenuButton struct {
	ID         string `orm:"ID" json:"ID"`                 // ID
	RoleId     string `orm:"RoleId" json:"RoleId"`         // 权限ID
	MenuId     string `orm:"MenuId" json:"MenuId"`         // 菜单ID
	ButtonId   string `orm:"ButtonId" json:"ButtonId"`     // 按钮ID
	CreateBy   string `orm:"CreateBy" json:"CreateBy"`     // 创建人
	CreateTime string `orm:"CreateTime" json:"CreateTime"` // 创建时间
	UpdateBy   string `orm:"UpdateBy" json:"UpdateBy"`     // 修改人
	UpdateTime string `orm:"UpdateTime" json:"UpdateTime"` // 修改时间
}

func (*RoleMenuButton) TableName() string {
	return "RoleMenuButton"
}

type UserView struct {
	ID             string `orm:"ID" json:"ID"`                   // 主键
	AccountName    string `orm:"AccountName" json:"AccountName"` // 帐户名
	PassWord       string `orm:"PassWord" json:"PassWord"`       // 帐户密码（32位MD5加密）
	RealName       string `orm:"RealName" json:"RealName"`       // 别名
	MobilePhone    string `orm:"MobilePhone" json:"MobilePhone"` // 联系方式
	Email          string `orm:"Email" json:"Email"`             // 邮箱
	Description    string `orm:"Description" json:"Description"` // 简介
	CreateBy       string `orm:"CreateBy" json:"CreateBy"`       // 创建人
	CreateTime     string `orm:"CreateTime" json:"CreateTime"`   // 创建时间
	UpdateBy       string `orm:"UpdateBy" json:"UpdateBy"`       // 修改人
	UpdateTime     string `orm:"UpdateTime" json:"UpdateTime"`   // 修改时间
	IsAble         int    `orm:"IsAble" json:"IsAble"`           // 是否有效
	IfChangePwd    int    `orm:"IfChangePwd" json:"IfChangePwd"` // 是否修改密码
	DepartmentName string `orm:"DepartmentName" json:"DepartmentName"`
	RoleName       string `orm:"RoleName" json:"RoleName"`
}

func (*UserView) TableName() string {
	return "UserView"
}

type FlowStep struct {
	ID         string `orm:"ID" json:"ID"`                 // 主键
	FlowId     string `orm:"FlowId" json:"FlowId"`         // 流程ID
	StepNum    int    `orm:"StepNum" json:"StepNum"`       // 审批步骤号
	StepType   int    `orm:"StepType" json:"StepType"`     // 审批方式（一人审核通过即可/全部审批通过/半数通过）
	Notes      string `orm:"Notes" json:"Notes"`           // 备注
	CreateBy   string `orm:"CreateBy" json:"CreateBy"`     // 创建人
	CreateTime string `orm:"CreateTime" json:"CreateTime"` // 创建时间
	UpdateBy   string `orm:"UpdateBy" json:"UpdateBy"`     // 修改人
	UpdateTime string `orm:"UpdateTime" json:"UpdateTime"` // 修改时间
}

func (*FlowStep) TableName() string {
	return "FlowStep"
}

type Department struct {
	ID          string `orm:"ID" json:"ID"`                   // 部门ID
	Name        string `orm:"Name" json:"Name"`               // 部门名
	ParentId    string `orm:"ParentId" json:"ParentId"`       // 父级部门名
	Sort        int    `orm:"Sort" json:"Sort"`               // 排序
	Description string `orm:"Description" json:"Description"` // 简介
	CreateBy    string `orm:"CreateBy" json:"CreateBy"`       // 创建人
	CreateTime  string `orm:"CreateTime" json:"CreateTime"`   // 创建时间
	UpdateBy    string `orm:"UpdateBy" json:"UpdateBy"`       // 修改人
	UpdateTime  string `orm:"UpdateTime" json:"UpdateTime"`   // 修改时间
	IsAble      int    `orm:"IsAble" json:"IsAble"`           // 是否有效
}

func (*Department) TableName() string {
	return "Department"
}

type Flow struct {
	ID          string `orm:"ID" json:"ID"`                   // 主键
	No          string `orm:"No" json:"No"`                   // 单据号
	MenuId      string `orm:"MenuId" json:"MenuId"`           // 菜单ID
	Description string `orm:"Description" json:"Description"` // 简介
	Notes       string `orm:"Notes" json:"Notes"`             // 备注
	CreateBy    string `orm:"CreateBy" json:"CreateBy"`       // 创建人
	CreateTime  string `orm:"CreateTime" json:"CreateTime"`   // 创建时间
	UpdateBy    string `orm:"UpdateBy" json:"UpdateBy"`       // 修改人
	UpdateTime  string `orm:"UpdateTime" json:"UpdateTime"`   // 修改时间
	Status      int    `orm:"Status" json:"Status"`           // 状态
}

func (*Flow) TableName() string {
	return "Flow"
}

type FlowStepLog struct {
	ID         string `orm:"ID" json:"ID"`                 // 主键
	FlowLogId  string `orm:"FlowLogId" json:"FlowLogId"`   // 流程日志ID
	FlowId     string `orm:"FlowId" json:"FlowId"`         // 流程ID
	StepId     string `orm:"StepId" json:"StepId"`         // 步骤ID
	BillId     string `orm:"BillId" json:"BillId"`         // 单据ID
	IsApproval int    `orm:"IsApproval" json:"IsApproval"` // 是否审批通过
	Notes      string `orm:"Notes" json:"Notes"`           // 备注
	CreateBy   string `orm:"CreateBy" json:"CreateBy"`     // 创建人
	CreateTime string `orm:"CreateTime" json:"CreateTime"` // 创建时间
	UpdateBy   string `orm:"UpdateBy" json:"UpdateBy"`     // 修改人
	UpdateTime string `orm:"UpdateTime" json:"UpdateTime"` // 修改时间
}

func (*FlowStepLog) TableName() string {
	return "FlowStepLog"
}

type MenuButton struct {
	ID         string `orm:"ID" json:"ID"`                 // ID
	MenuId     string `orm:"MenuId" json:"MenuId"`         // 菜单ID
	ButtonId   string `orm:"ButtonId" json:"ButtonId"`     // 按钮ID
	CreateBy   string `orm:"CreateBy" json:"CreateBy"`     // 创建人
	CreateTime string `orm:"CreateTime" json:"CreateTime"` // 创建时间
	UpdateBy   string `orm:"UpdateBy" json:"UpdateBy"`     // 修改人
	UpdateTime string `orm:"UpdateTime" json:"UpdateTime"` // 修改时间
}

func (*MenuButton) TableName() string {
	return "MenuButton"
}
