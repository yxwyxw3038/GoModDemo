package consts

var MsgFlags = map[int]string{
	SUCCESS: "1",
	ERROR:   "-1",
	None:    "0",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
