package util
import ("strconv")
func  ToString (t interface {} ) string {
	if t==nil {
		return ""
	}
	s:=t.(string)
	return s
}

func  ToInt (t interface {} ) int {
	if t==nil {
		return 0
	}
	s:=0
	switch v := t.(type) {
	case int:
		s=t.(int)
	// case int32:
	// 	strInt32 := strconv.FormatInt(v, 10)
	// 	s ,_ := strconv.Atoi(strInt32)
	case int64:
		strInt64 := strconv.FormatInt(v, 10)
		s ,_ = strconv.Atoi(strInt64)
	case string:
		s ,_ = strconv.Atoi(v)
	default:
		s=0
	}	
	return s
}