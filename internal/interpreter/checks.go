package interpreter

func checkValueType[T any](value interface{}) bool {
	_, ok := value.(T)
	return ok
}
