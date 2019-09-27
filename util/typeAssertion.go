package util

func  ToString (t interface {} ) string {
	if t==nil {
		return ""
	}
	s:=t.(string)
	return s
}