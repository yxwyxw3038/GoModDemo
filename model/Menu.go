package model
type MenuTree struct {
	ID          string `:"ID"`          
	Icon string `:"Icon"` 
	Url    string `:"Url"`    
	ParentId    string `:"ParentId"`    
	MenuName       string `:"MenuName"`       
	Node []MenuTree
}

type CascaderMenu struct {
	Value          string `json:"value"`          
	Label string `json:"label"`   
	Children []CascaderMenu  `json:"children"` 
}