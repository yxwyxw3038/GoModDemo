package consts

func GetUserTabInfo() []string {

	info := []string{"Description", "IfChangePwd", "IsAble"}
	return info
}

func GetMenuTabInfo() []string {

	info := []string{"Description", "IsAble"}
	return info
}
func GetDeptInfo() []string {

	info := []string{"Description", "ParentId", "IsAble", "CreateBy", "UpdateBy"}
	return info
}
func GetTabInfo() []string {

	info := []string{"Description", "IsAble", "CreateBy", "UpdateBy"}
	return info
}

func GetUserTokenInfo() []string {

	info := []string{"IsLoginOut", "Port", "Address", "CreateTime", "UpdateTime", "LoginOutTime"}
	return info
}
