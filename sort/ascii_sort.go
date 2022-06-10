package sort

import "sort"

// AscAscii
// 对map的key的首字母按Ascii码升序排序
func AscAscii(params map[string]interface{}) []string {
	if len(params) == 0 {
		panic("params can't be empty!")
	}

	var temp []string

	for k := range params {
		temp = append(temp, k)
	}

	sort.Strings(temp)

	return temp
}

// DescAscii
// 对map的key的首字母按Ascii码降序排序
func DescAscii(params map[string]interface{}) []string {
	count := len(params)
	if count == 0 {
		panic("params can't be empty!")
	}

	var temp []string

	for k := range params {
		temp = append(temp, k)
	}

	sort.Strings(temp)

	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}

	return temp
}