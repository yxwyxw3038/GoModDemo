package consts

func GetTabInfo() []string {

	info := []string{"Description", "IsAble", "CreateBy", "UpdateBy", "CreateTime", "UpdateTime"}
	return info
}
func GetUserTabInfo() []string {

	info := []string{"IfChangePwd"}
	return append(info, GetTabInfo()...)
}

func GetMenuTabInfo() []string {

	return GetTabInfo()
}
func GetDeptInfo() []string {

	info := []string{"ParentId"}
	return append(info, GetTabInfo()...)
}

func GetUserTokenInfo() []string {

	info := []string{"IsLoginOut", "Port", "Address", "CreateTime", "UpdateTime", "LoginOutTime"}
	return info
}

func GetBillNoInfo() []string {

	info := []string{"CurrentId", "CurrentBillNo", "CurrentTime", "MaskInfo", "EndLength"}
	return append(info, GetTabInfo()...)
}
func GetNoticeInfo() []string {

	info := []string{"TypeId", "Title", "Content", "Notes", "Status", "NoticeTime", "SendBeginTime", "SendEndTime"}
	return append(info, GetTabInfo()...)
}
