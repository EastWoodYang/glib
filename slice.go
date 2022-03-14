package glib

import (
	"fmt"
	"strconv"
	"strings"
)

/* ================================================================================
 * 常用
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */


/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取指定个数的uint64slice
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetUint64SliceRange(int64Slice []uint64, count int) []uint64 {
	length := len(int64Slice)
	if length > count-1 {
		length = count
	}
	return int64Slice[0:length]
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 翻转uint64 Slice
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func ReverseUint64Slice(int64Slice []uint64) {
	for from, to := 0, len(int64Slice)-1; from < to; from, to = from+1, to-1 {
		int64Slice[from], int64Slice[to] = int64Slice[to], int64Slice[from]
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 用指定的字符串把源字符串分隔为uint64切片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringToUint64Slice(sourceString string, args ...string) []uint64 {
	result := make([]uint64, 0)

	if len(sourceString) == 0 {
		return result
	}

	splitString := ","
	if len(args) == 1 {
		splitString = args[0]
	}

	stringSlice := strings.Split(sourceString, splitString)
	for _, v := range stringSlice {
		if value, err := strconv.ParseUint(v, 10, 64); err == nil {
			result = append(result, value)
		}
	}

	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 用指定的字符串把源字符串分隔为int64切片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringToInt64Slice(sourceString string, args ...string) []int64 {
	result := make([]int64, 0)

	if len(sourceString) == 0 {
		return result
	}

	splitString := ","
	if len(args) == 1 {
		splitString = args[0]
	}

	stringSlice := strings.Split(sourceString, splitString)
	for _, v := range stringSlice {
		if value, err := strconv.ParseInt(v, 10, 64); err == nil {
			result = append(result, value)
		}
	}

	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 用指定的字符串把源字符串分隔为int切片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringToIntSlice(sourceString string, args ...string) []int {
	result := make([]int, 0)

	if len(sourceString) == 0 {
		return result
	}

	splitString := ","
	if len(args) == 1 {
		splitString = args[0]
	}

	stringSlice := strings.Split(sourceString, splitString)
	for _, v := range stringSlice {
		if value, err := strconv.Atoi(v); err == nil {
			result = append(result, value)
		}
	}

	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符串切片转为整型64切片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringSliceToInt64Slice(values []string) []int64 {
	results := make([]int64, 0)

	for _, value := range values {
		valueInt64 := StringToInt64(value)
		results = append(results, valueInt64)
	}

	return results
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符串切片转为无符号整型64切片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringSliceToUint64Slice(values []string) []uint64 {
	results := make([]uint64, 0)

	for _, value := range values {
		valueInt64 := StringToUint64(value)
		results = append(results, valueInt64)
	}

	return results
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 整型64切片转为字符串切片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Int64SliceToStringSlice(values []int64) []string {
	results := make([]string, 0)

	for _, value := range values {
		valueString := Int64ToString(value)
		results = append(results, valueString)
	}

	return results
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 无符号整型64切片转为字符串切片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Uint64SliceToStringSlice(values []uint64) []string {
	results := make([]string, 0)

	for _, value := range values {
		valueString := Uint64ToString(value)
		results = append(results, valueString)
	}

	return results
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 用指定的字符串把uint64切片链接为字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Uint64SliceToString(uintSlice []uint64, args ...string) string {
	result := ""

	if len(uintSlice) == 0 {
		return result
	}

	joinString := ","
	if len(args) == 1 {
		joinString = args[0]
	}

	count := len(uintSlice)
	if count == 1 {
		result = fmt.Sprintf("%d", uintSlice[0])
	} else if count > 1 {
		for _, value := range uintSlice {
			valueString := fmt.Sprintf("%d", value)
			if len(result) == 0 {
				result = result + valueString
			} else {
				result = result + joinString + valueString
			}
		}
	}

	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 用指定的字符串把int64切片链接为字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Int64SliceToString(intSlice []int64, args ...string) string {
	result := ""

	if len(intSlice) == 0 {
		return result
	}

	joinString := ","
	if len(args) == 1 {
		joinString = args[0]
	}

	count := len(intSlice)
	if count == 1 {
		result = fmt.Sprintf("%d", intSlice[0])
	} else if count > 1 {
		for _, value := range intSlice {
			valueString := fmt.Sprintf("%d", value)
			if len(result) == 0 {
				result = result + valueString
			} else {
				result = result + joinString + valueString
			}
		}
	}

	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 用指定的字符串分隔源字符串为字符串切片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringToStringSlice(sourceString string, args ...string) []string {
	result := make([]string, 0)

	if len(sourceString) == 0 {
		return result
	}

	splitString := ","
	if len(args) == 1 {
		splitString = args[0]
	}

	stringSlice := strings.Split(sourceString, splitString)
	for _, v := range stringSlice {
		if v != "" {
			result = append(result, v)
		}
	}

	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 用指定的字符串链接字符串切片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringSliceToString(stringSlice []string, args ...string) string {
	result := ""

	if len(stringSlice) == 0 {
		return result
	}

	joinString := ","
	if len(args) == 1 {
		joinString = args[0]
	}

	if len(stringSlice) == 1 {
		result = strings.Join(stringSlice, "")
	} else {
		for _, v := range stringSlice {
			if len(result) == 0 {
				result = result + v
			} else {
				result = result + joinString + v
			}
		}
	}

	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 保留指定长度字符串切片，前面的数据移除
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringSliceLatest(srcSlice []string, maxCount int) []string {
	destSlice := srcSlice
	count := len(destSlice)
	if count > maxCount {
		offsetIndex := count - maxCount
		destSlice = destSlice[offsetIndex:count]
	}

	return destSlice
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断字符串切片及单个项的字符数是否匹配指定大小，
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func IsMatchStringSliceCount(srcSlice []string, maxCount, stringItemCount int) bool {
	isMatch := true
	srcSliceCount := len(srcSlice)

	if srcSliceCount > 0 {
		if srcSliceCount > maxCount {
			isMatch = false
		}

		for _, stringItem := range srcSlice {
			if GetStringCount(stringItem) > stringItemCount {
				isMatch = false
				break
			}
		}
	}

	return isMatch
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 过滤int64数组（从all中过滤所有other中的数据，返回未被过滤的数据集合）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func FilterInt64Slice(all, other []int64) []int64 {
	allMaps := make(map[int64]bool, 0)
	sliceResult := make([]int64, 0)

	for _, v := range all {
		allMaps[v] = true
	}

	for _, v := range other {
		if _, isOk := allMaps[v]; isOk {
			allMaps[v] = false
		}
	}

	for _, v := range all {
		if allMaps[v] {
			sliceResult = append(sliceResult, v)
		}
	}

	return sliceResult
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 过滤uint64数组（从all中过滤所有other中的数据，返回未被过滤的数据集合）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func FilterUint64Slice(all, other []uint64) []uint64 {
	allMaps := make(map[uint64]bool, 0)
	sliceResult := make([]uint64, 0)

	for _, v := range all {
		allMaps[v] = true
	}

	for _, v := range other {
		if _, isOk := allMaps[v]; isOk {
			allMaps[v] = false
		}
	}

	for _, v := range all {
		if allMaps[v] {
			sliceResult = append(sliceResult, v)
		}
	}

	return sliceResult
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 过滤字符数组（从all中过滤所有other中的数据，返回未被过滤的数据集合）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func FilterStringSlice(all, other []string) []string {
	allMaps := make(map[string]bool, 0)
	sliceResult := make([]string, 0)

	for _, v := range all {
		allMaps[v] = true
	}

	for _, v := range other {
		if _, isOk := allMaps[v]; isOk {
			allMaps[v] = false
		}
	}

	for _, v := range all {
		if allMaps[v] {
			sliceResult = append(sliceResult, v)
		}
	}

	return sliceResult
}
