package model

type MenuTree struct {
	ID       string `:"ID"`
	Icon     string `:"Icon"`
	Url      string `:"Url"`
	ParentId string `:"ParentId"`
	MenuName string `:"MenuName"`
	Node     []MenuTree
}

type CascaderMenu struct {
	Value    string         `json:"value"`
	Label    string         `json:"label"`
	Children []CascaderMenu `json:"children"`
}

type MenuView struct {
	ID          string `:"ID"`          // 主键
	Name        string `:"Name"`        // 菜单名
	ParentId    string `:"ParentId"`    // 父级菜单ID
	ParentName  string `:"ParentName"`  // 父级菜单名
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


