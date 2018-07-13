package jx

func asInt32(n interface{}) int32 {
	i, ok := n.(int32)
	if ok {
		return i
	}
	i2, ok := n.(int)
	if ok {
		return int32(i2)
	}
	return 0
}

func asString(n interface{}) string {
	s, ok := n.(string)
	if ok {
		return s
	}
	return ""
}

func asBool(n interface{}) bool {
	b, ok := n.(bool)
	if ok {
		return b
	}
	return asString(n) == "true"
}
