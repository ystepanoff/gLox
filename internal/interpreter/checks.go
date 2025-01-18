package interpreter

func checkValueType[T any](value interface{}) bool {
	_, ok := value.(T)
	return ok
}

func checkValuesType[T any](args ...interface{}) bool {
	for _, value := range args {
		if !checkValueType[T](value) {
			return false
		}
	}
	return true
}
