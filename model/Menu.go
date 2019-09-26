package model
type MenuTree struct {
	ID          string `:"ID"`          
	Icon string `:"Icon"` 
	Url    string `:"Url"`    
	ParentId    string `:"ParentId"`    
	MobilePhone string `:"MobilePhone"` 
	MenuName       string `:"MenuName"`       
	Node []MenuTree
}