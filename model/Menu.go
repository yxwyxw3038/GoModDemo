package model
type MenuTree struct {
	ID          string `:"ID"`          
	Icon string `:"Icon"` 
	Url    string `:"Url"`    
	ParentId    string `:"ParentId"`    
	MenuName       string `:"MenuName"`       
	Node []MenuTree
}