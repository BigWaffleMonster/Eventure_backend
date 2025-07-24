package helpers

func MapArray[T any, U any](arr *[]T, fn func(T) U) *[]U {
	result := make([]U, len(*arr))
	for i, item := range *arr {
		result[i] = fn(item)
	}
	return &result
}