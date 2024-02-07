package tool

import "reflect"

func DeduplicateSlice(slice interface{}) interface{} {
	// 获取切片的反射值
	sliceValue := reflect.ValueOf(slice)

	// 如果传入的不是切片类型，则直接返回
	if sliceValue.Kind() != reflect.Slice {
		return slice
	}

	// 创建一个 map 用于存储切片中的元素，以及一个用于存储去重后的切片
	uniqueMap := make(map[interface{}]bool)
	uniqueSlice := reflect.MakeSlice(sliceValue.Type(), 0, sliceValue.Len())

	// 遍历切片，将元素添加到 map 中进行去重
	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i).Interface()
		if !uniqueMap[element] {
			uniqueMap[element] = true
			uniqueSlice = reflect.Append(uniqueSlice, reflect.ValueOf(element))
		}
	}

	// 将 map 中的元素拷贝到新的切片中，并返回
	return uniqueSlice.Interface()
}
