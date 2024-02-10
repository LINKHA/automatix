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

// RemoveItem 从切片中删除指定的元素
func RemoveItemSlice[T comparable](slice []T, item T) []T {
	// 创建一个新切片用于存储结果
	result := make([]T, 0)

	// 遍历原始切片
	for _, v := range slice {
		// 如果元素不等于要删除的元素，则添加到结果切片中
		if v != item {
			result = append(result, v)
		}
	}

	return result
}
