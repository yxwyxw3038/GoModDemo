package model

type MenuTree struct {
	ID       string `:"ID"`
	Icon     string `:"Icon"`
	Url      string `:"Url"`
	ParentId string `:"ParentId"`
	MenuName string `:"MenuName"`
	Node     []MenuTree
}

type CascaderModel struct {
	Value    string          `json:"value"`
	Label    string          `json:"label"`
	Children []CascaderModel `json:"children"`
}

type MenuView struct {
	Menu
	ParentName string `:"ParentName"` // 父级菜单名

}
