package consts
var MsgFlags = map[int]string{
    SUCCESS:                         "ok",
    ERROR:                           "fail",
    None:                  "none",
}

func GetMsg(code int) string {
    msg, ok := MsgFlags[code]
    if ok {
        return msg
    }
    return MsgFlags[ERROR]
}