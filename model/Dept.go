package model

type DeptView struct {
	Department
	ParentName string `:"ParentName"` // 父级部门名

}
