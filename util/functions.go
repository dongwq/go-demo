package util

// the func is the same as condition ? true : false
func Ternary(express bool, trueVal interface{}, falseVal interface{}) interface{} {
	if express {
		return trueVal
	}
	return falseVal
}
