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

type MenuButton struct {
	ID         string `:"ID"`         // ID
	MenuId     string `:"MenuId"`     // 菜单ID
	ButtonId   string `:"ButtonId"`   // 按钮ID
	CreateBy   string `:"CreateBy"`   // 创建人
	CreateTime string `:"CreateTime"` // 创建时间
	UpdateBy   string `:"UpdateBy"`   // 修改人
	UpdateTime string `:"UpdateTime"` // 修改时间
}

type Notice struct {
	ID            string `:"ID"`            // 主键
	No            string `:"No"`            // 通知单号
	TypeId        int    `:"TypeId"`        // 通知类型
	Title         string `:"Title"`         // 通知标题
	Content       string `:"Content"`       // 通知内容
	Notes         string `:"Notes"`         // 备注
	CreateBy      string `:"CreateBy"`      // 创建人
	CreateTime    string `:"CreateTime"`    // 创建时间
	UpdateBy      string `:"UpdateBy"`      // 修改人
	UpdateTime    string `:"UpdateTime"`    // 修改时间
	Status        int    `:"Status"`        // 通知状态
	NoticeTime    string `:"NoticeTime"`    // 通知时间
	SendBeginTime string `:"SendBeginTime"` // 通知发送开始时间
	SendEndTime   string `:"SendEndTime"`   // 通知发送结束时间
}

type Department struct {
	ID          string `:"ID"`          // 部门ID
	Name        string `:"Name"`        // 部门名
	ParentId    string `:"ParentId"`    // 父级部门名
	Sort        int    `:"Sort"`        // 排序
	Description string `:"Description"` // 简介
	CreateBy    string `:"CreateBy"`    // 创建人
	CreateTime  string `:"CreateTime"`  // 创建时间
	UpdateBy    string `:"UpdateBy"`    // 修改人
	UpdateTime  string `:"UpdateTime"`  // 修改时间
	IsAble      int    `:"IsAble"`      // 是否有效
}

type DeptView struct {
	ID          string `:"ID"`          // 部门ID
	Name        string `:"Name"`        // 部门名
	ParentId    string `:"ParentId"`    // 父级部门名
	Sort        int    `:"Sort"`        // 排序
	Description string `:"Description"` // 简介
	CreateBy    string `:"CreateBy"`    // 创建人
	CreateTime  string `:"CreateTime"`  // 创建时间
	UpdateBy    string `:"UpdateBy"`    // 修改人
	UpdateTime  string `:"UpdateTime"`  // 修改时间
	IsAble      int    `:"IsAble"`      // 是否有效
	ParentName  string `:"ParentName"`
}

type Flow struct {
	ID          string `:"ID"`          // 主键
	No          string `:"No"`          // 单据号
	MenuId      string `:"MenuId"`      // 菜单ID
	Description string `:"Description"` // 简介
	Notes       string `:"Notes"`       // 备注
	CreateBy    string `:"CreateBy"`    // 创建人
	CreateTime  string `:"CreateTime"`  // 创建时间
	UpdateBy    string `:"UpdateBy"`    // 修改人
	UpdateTime  string `:"UpdateTime"`  // 修改时间
	Status      int    `:"Status"`      // 状态
}
type FlowView struct {
	ID          string `:"ID"`          // 主键
	No          string `:"No"`          // 单据号
	MenuId      string `:"MenuId"`      // 菜单ID
	Description string `:"Description"` // 简介
	Notes       string `:"Notes"`       // 备注
	CreateBy    string `:"CreateBy"`    // 创建人
	CreateTime  string `:"CreateTime"`  // 创建时间
	UpdateBy    string `:"UpdateBy"`    // 修改人
	UpdateTime  string `:"UpdateTime"`  // 修改时间
	Status      int    `:"Status"`      // 状态
	MenuCode    string `:"MenuCode"`    // 菜单代码
	MenuName    string `:"MenuName"`    // 菜单名称
	StepNum     int    `:"StepNum"`     // 步骤数
}

type FlowLog struct {
	ID          string `:"ID"`          // 主键
	FlowId      string `:"FlowId"`      // 流程ID
	BillId      string `:"BillId"`      // 单据ID
	IsApproval  int    `:"IsApproval"`  // 是否审批通过
	Description string `:"Description"` // 简介
	Notes       string `:"Notes"`       // 备注
	CreateBy    string `:"CreateBy"`    // 创建人
	CreateTime  string `:"CreateTime"`  // 创建时间
	UpdateBy    string `:"UpdateBy"`    // 修改人
	UpdateTime  string `:"UpdateTime"`  // 修改时间
	Status      int    `:"Status"`      // 状态
}

type FlowStepUserLog struct {
	ID         string `:"ID"`         // 主键
	FlowLogId  string `:"FlowLogId"`  // 流程日志ID
	StepLogId  string `:"StepLogId"`  // 步骤日志ID
	FlowId     string `:"FlowId"`     // 流程ID
	StepId     string `:"StepId"`     // 步骤ID
	UserId     string `:"UserId"`     // 用户ID
	Notes      string `:"Notes"`      // 备注
	CreateBy   string `:"CreateBy"`   // 创建人
	CreateTime string `:"CreateTime"` // 创建时间
	UpdateBy   string `:"UpdateBy"`   // 修改人
	UpdateTime string `:"UpdateTime"` // 修改时间
	IsApproval int    `:"IsApproval"` // 是否审批通过
}

type Icons struct {
	ID         string `:"ID"`         // ID
	Code       string `:"Code"`       // 图标代码
	CodeName   string `:"CodeName"`   // 图标名称
	CssInfo    string `:"CssInfo"`    // 图标样式
	CreateBy   string `:"CreateBy"`   // 创建人
	CreateTime string `:"CreateTime"` // 创建时间
	UpdateBy   string `:"UpdateBy"`   // 修改人
	UpdateTime string `:"UpdateTime"` // 修改时间
}

type Parameter struct {
	ID          string `:"ID"`          // 主键
	Name        string `:"Name"`        // 参数名称
	Code        string `:"Code"`        // 参数代码
	ParentId    string `:"ParentId"`    // 上级参数
	Sort        int    `:"Sort"`        // 排序
	Description string `:"Description"` // 简介
	CreateBy    string `:"CreateBy"`    // 创建人
	CreateTime  string `:"CreateTime"`  // 创建时间
	UpdateBy    string `:"UpdateBy"`    // 修改人
	UpdateTime  string `:"UpdateTime"`  // 修改时间
	IsAble      int    `:"IsAble"`      // 是否有效
}

type RoleMenu struct {
	ID         string `:"ID"`         // ID
	RoleId     string `:"RoleId"`     // 权限ID
	MenuId     string `:"MenuId"`     // 菜单ID
	CreateBy   string `:"CreateBy"`   // 创建人
	CreateTime string `:"CreateTime"` // 创建时间
	UpdateBy   string `:"UpdateBy"`   // 修改人
	UpdateTime string `:"UpdateTime"` // 修改时间
}

type RoleMenuButton struct {
	ID         string `:"ID"`         // ID
	RoleId     string `:"RoleId"`     // 权限ID
	MenuId     string `:"MenuId"`     // 菜单ID
	ButtonId   string `:"ButtonId"`   // 按钮ID
	CreateBy   string `:"CreateBy"`   // 创建人
	CreateTime string `:"CreateTime"` // 创建时间
	UpdateBy   string `:"UpdateBy"`   // 修改人
	UpdateTime string `:"UpdateTime"` // 修改时间
}

type UserDepartment struct {
	ID           string `:"ID"`           // 部门ID
	UserId       string `:"UserId"`       // 用户ID
	DepartmentId string `:"DepartmentId"` // 部门ID
	CreateBy     string `:"CreateBy"`     // 创建人
	CreateTime   string `:"CreateTime"`   // 创建时间
	UpdateBy     string `:"UpdateBy"`     // 修改人
	UpdateTime   string `:"UpdateTime"`   // 修改时间
}

type UserRole struct {
	ID         string `:"ID"`         // 部门ID
	UserId     string `:"UserId"`     // 用户ID
	RoleId     string `:"RoleId"`     // 权限ID
	CreateBy   string `:"CreateBy"`   // 创建人
	CreateTime string `:"CreateTime"` // 创建时间
	UpdateBy   string `:"UpdateBy"`   // 修改人
	UpdateTime string `:"UpdateTime"` // 修改时间
}

type UserToken struct {
	Token        string `:"Token"`        // 主键
	UserId       string `:"UserId"`       // 用户ID
	Port         string `:"Port"`         // 端口
	Address      string `:"Address"`      // 地址
	CreateTime   string `:"CreateTime"`   // 创建时间
	UpdateTime   string `:"UpdateTime"`   // 修改时间
	IsLoginOut   int    `:"IsLoginOut"`   // 离线标记
	LoginOutTime string `:"LoginOutTime"` // 离线时间
}

type FlowStepLog struct {
	ID         string `:"ID"`         // 主键
	FlowLogId  string `:"FlowLogId"`  // 流程日志ID
	FlowId     string `:"FlowId"`     // 流程ID
	StepId     string `:"StepId"`     // 步骤ID
	BillId     string `:"BillId"`     // 单据ID
	IsApproval int    `:"IsApproval"` // 是否审批通过
	Notes      string `:"Notes"`      // 备注
	CreateBy   string `:"CreateBy"`   // 创建人
	CreateTime string `:"CreateTime"` // 创建时间
	UpdateBy   string `:"UpdateBy"`   // 修改人
	UpdateTime string `:"UpdateTime"` // 修改时间
}

type FlowStepUser struct {
	ID         string `:"ID"`         // 主键
	FlowId     string `:"FlowId"`     // 流程ID
	StepId     string `:"StepId"`     // 步骤ID
	UserId     string `:"UserId"`     // 用户ID
	Notes      string `:"Notes"`      // 备注
	CreateBy   string `:"CreateBy"`   // 创建人
	CreateTime string `:"CreateTime"` // 创建时间
	UpdateBy   string `:"UpdateBy"`   // 修改人
	UpdateTime string `:"UpdateTime"` // 修改时间
}
type FlowStepUserView struct {
	ID          string `:"ID"`          // 主键
	FlowId      string `:"FlowId"`      // 流程ID
	StepId      string `:"StepId"`      // 步骤ID
	UserId      string `:"UserId"`      // 用户ID
	Notes       string `:"Notes"`       // 备注
	CreateBy    string `:"CreateBy"`    // 创建人
	CreateTime  string `:"CreateTime"`  // 创建时间
	UpdateBy    string `:"UpdateBy"`    // 修改人
	UpdateTime  string `:"UpdateTime"`  // 修改时间
	AccountName string `:"AccountName"` // 帐号
	RealName    string `:"RealName"`    // 用户名称
}
type BillNo struct {
	ID            string `:"ID"`            // ID
	Code          string `:"Code"`          // 单位据代码
	CodeName      string `:"CodeName"`      // 单据号名称
	CurrentId     int    `:"CurrentId"`     // 当前流水号
	CurrentBillNo string `:"CurrentBillNo"` // 当前单据号
	CurrentTime   string `:"CurrentTime"`   // 当前单据生成时间
	MaskInfo      string `:"MaskInfo"`      // 掩码
	EndLength     int    `:"EndLength"`     // 流水号长度
	CreateBy      string `:"CreateBy"`      // 创建人
	CreateTime    string `:"CreateTime"`    // 创建时间
	UpdateBy      string `:"UpdateBy"`      // 修改人
	UpdateTime    string `:"UpdateTime"`    // 修改时间
}

type Button struct {
	ID          string `:"ID"`          // 主键
	Name        string `:"Name"`        // 按钮名
	Code        string `:"Code"`        // 按钮代码
	Icon        string `:"Icon"`        // 按钮图标
	Sort        int    `:"Sort"`        // 排序
	Description string `:"Description"` // 简介
	CreateBy    string `:"CreateBy"`    // 创建人
	CreateTime  string `:"CreateTime"`  // 创建时间
	UpdateBy    string `:"UpdateBy"`    // 修改人
	UpdateTime  string `:"UpdateTime"`  // 修改时间
	IsAble      int    `:"IsAble"`      // 是否有效
}

type Menu struct {
	ID          string `:"ID"`          // 主键
	Name        string `:"Name"`        // 菜单名
	ParentId    string `:"ParentId"`    // 父级菜单名
	Code        string `:"Code"`        // 菜单代码
	LinkAddress string `:"LinkAddress"` // 菜单地址
	Icon        string `:"Icon"`        // 菜单图标
	Sort        int    `:"Sort"`        // 排序
	Description string `:"Description"` // 简介
	CreateBy    string `:"CreateBy"`    // 创建人
	CreateTime  string `:"CreateTime"`  // 创建时间
	UpdateBy    string `:"UpdateBy"`    // 修改人
	UpdateTime  string `:"UpdateTime"`  // 修改时间
	IsAble      int    `:"IsAble"`      // 是否有效
}

type NoticeUser struct {
	ID         string `:"ID"`         // 主键
	NoticeId   string `:"NoticeId"`   // 通知单ID
	UserId     string `:"UserId"`     // 用户ID
	Notes      string `:"Notes"`      // 备注
	CreateBy   string `:"CreateBy"`   // 创建人
	CreateTime string `:"CreateTime"` // 创建时间
	UpdateBy   string `:"UpdateBy"`   // 修改人
	UpdateTime string `:"UpdateTime"` // 修改时间
	SendFlag   int    `:"SendFlag"`   // 发送标记
	SendTime   string `:"SendTime"`   // 发送时间
}

type Role struct {
	ID          string `:"ID"`          // 主键
	Name        string `:"Name"`        // 权限名
	Description string `:"Description"` // 简介
	CreateBy    string `:"CreateBy"`    // 创建人
	CreateTime  string `:"CreateTime"`  // 创建时间
	UpdateBy    string `:"UpdateBy"`    // 修改人
	UpdateTime  string `:"UpdateTime"`  // 修改时间
	IsAble      int    `:"IsAble"`      // 是否有效
}

type User struct {
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
}

type FlowStep struct {
	ID         string `:"ID"`         // 主键
	FlowId     string `:"FlowId"`     // 流程ID
	StepNum    int    `:"StepNum"`    // 审批步骤号
	StepType   int    `:"StepType"`   // 审批方式（一人审核通过即可/全部审批通过/半数通过）
	Notes      string `:"Notes"`      // 备注
	CreateBy   string `:"CreateBy"`   // 创建人
	CreateTime string `:"CreateTime"` // 创建时间
	UpdateBy   string `:"UpdateBy"`   // 修改人
	UpdateTime string `:"UpdateTime"` // 修改时间
}

type MenuView struct {
	ID          string `:"ID"`          // 主键
	Name        string `:"Name"`        // 菜单名
	ParentId    string `:"ParentId"`    // 父级菜单名
	Code        string `:"Code"`        // 菜单代码
	LinkAddress string `:"LinkAddress"` // 菜单地址
	Icon        string `:"Icon"`        // 菜单图标
	Sort        int    `:"Sort"`        // 排序
	Description string `:"Description"` // 简介
	CreateBy    string `:"CreateBy"`    // 创建人
	CreateTime  string `:"CreateTime"`  // 创建时间
	UpdateBy    string `:"UpdateBy"`    // 修改人
	UpdateTime  string `:"UpdateTime"`  // 修改时间
	IsAble      int    `:"IsAble"`      // 是否有效
	ParentName  string `:"ParentName"`
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
	RoleName       string `:"RoleName"`
}
