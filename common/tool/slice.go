package tool

// 使用模板函数实现切片去重
func DeduplicateSlice[T comparable](slice []T) []T {
	encountered := map[T]bool{}
	result := []T{}

	for _, v := range slice {
		if encountered[v] == false {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
}
