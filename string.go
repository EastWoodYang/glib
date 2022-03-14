package glib

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前日期流水字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetSeqNumber() string {
	currentDate := time.Now()
	format := "20060102150405.999999999"
	dateString := TimeToString(currentDate, format)
	dateString = strings.Replace(dateString, ".", "", -1)
	rndNumString := RandNumberString(6)
	return fmt.Sprintf("%s%s", dateString, rndNumString)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前唯一Token字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetTokenString() string {
	timestamp := UnixNanoTimestamp()
	return Md5(strconv.FormatInt(timestamp, 10))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断内容字符串头是否包含指定的字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HasPrefix(content, target string) bool {
	isPrefix := false

	if len(content) > 0 && len(target) > 0 {
		isPrefix = strings.HasPrefix(content, target)
	}

	return isPrefix
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断内容字符串尾是否包含指定的字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HasSuffix(content, target string) bool {
	isSuffix := false

	if len(content) > 0 && len(target) > 0 {
		isSuffix = strings.HasSuffix(content, target)
	}

	return isSuffix
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断内容字符串头尾是否包含指定的字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HasPrefixSuffix(content, target string) bool {
	isPrefixSuffix := false

	if len(content) > 0 && len(target) > 0 {
		isPrefix := strings.HasPrefix(content, target)
		isSuffix := strings.HasSuffix(content, target)

		isPrefixSuffix = isPrefix && isSuffix
	}

	return isPrefixSuffix
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符串替换
 * sourceString: 原始字符串
 * args[0...n-2]: 被替换字符串集合
 * args[n-1]: 替换字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringReplace(sourceString string, args ...string) string {
	target := ""
	replaces := []string{}
	result := sourceString
	count := len(args)

	if len(result) > 0 && count > 0 {
		if count == 1 {
			replaces = append(replaces, args[0])
		} else if count > 1 {
			for index, value := range args {
				if index == count-1 {
					target = value
				} else {
					replaces = append(replaces, value)
				}
			}
		}

		for _, value := range replaces {
			result = strings.Replace(result, value, target, -1)
		}
	}

	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取字符串个数（不是字节数）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetStringCount(sourceString string) int {
	if sourceString == "" {
		return 0
	}
	return utf8.RuneCountInString(sourceString)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取指定长度的字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetSubString(sourceString string, count int, args ...string) string {
	if len(sourceString) <= count {
		return sourceString
	}

	newString, sourceStringRune := "", []rune(sourceString)
	sl, rl := 0, 0

	more := ""
	isRune := true

	if len(args) > 0 {
		more = args[0]
	}

	for _, r := range sourceStringRune {
		if isRune {
			rl = 1
		} else {
			if int(r) < 128 {
				rl = 1
			} else {
				rl = 2
			}
		}

		if sl+rl > count {
			break
		}

		sl += rl

		newString += string(r)
	}

	if sl < len(sourceStringRune) {
		if len(more) > 0 {
			newString += more
		}
	}

	return newString
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 过滤主机协议
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func FilterHostProtocol(path string) string {
	if len(path) > 0 {
		path = strings.Trim(path, " ")

		if paths := StringToStringSlice(path, ":"); len(paths) > 1 {
			path = paths[1]
		}

		path = strings.TrimPrefix(path, "//")
		path = strings.TrimPrefix(path, "/")
		path = strings.TrimSuffix(path, "/")
	}

	return path
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 翻转字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func ReverseString(sourceString string) string {
	sourceRunes := []rune(sourceString)
	for from, to := 0, len(sourceRunes)-1; from < to; from, to = from+1, to-1 {
		sourceRunes[from], sourceRunes[to] = sourceRunes[to], sourceRunes[from]
	}
	return string(sourceRunes)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符串转换为bool
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringToBool(stringValue string) bool {
	var boolValue bool = true

	if len(stringValue) == 0 ||
		stringValue == "0" ||
		strings.ToLower(stringValue) == "f" ||
		strings.ToLower(stringValue) == "false" {

		boolValue = false
	}

	return boolValue
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符串转换为int32
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringToInt32(stringValue string) int32 {
	var intValue int64 = 0

	if len(stringValue) == 0 {
		return int32(intValue)
	}

	intValue, err := strconv.ParseInt(stringValue, 10, 32)
	if err != nil {
		return 0
	}

	return int32(intValue)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符串转换为uint32
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringToUint32(stringValue string) uint32 {
	var uintValue uint64 = 0

	if len(stringValue) == 0 {
		return uint32(uintValue)
	}

	uintValue, err := strconv.ParseUint(stringValue, 10, 32)
	if err != nil {
		return 0
	}

	return uint32(uintValue)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符串转换为uint64
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringToUint64(stringValue string) uint64 {
	var uintValue uint64 = 0

	if len(stringValue) == 0 {
		return uintValue
	}

	uintValue, err := strconv.ParseUint(stringValue, 10, 64)
	if err != nil {
		return 0
	}

	return uintValue
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符串转换为int64
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringToInt64(stringValue string) int64 {
	var intValue int64 = 0

	if len(stringValue) == 0 {
		return intValue
	}

	intValue, err := strconv.ParseInt(stringValue, 10, 64)
	if err != nil {
		return 0
	}

	return intValue
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符串转换为float64
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringToFloat64(stringValue string) float64 {
	var floatValue float64 = 0.0

	if len(stringValue) == 0 {
		return floatValue
	}

	floatValue, err := strconv.ParseFloat(stringValue, 64)
	if err != nil {
		return 0.0
	}

	return floatValue
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 无符号整型64转为字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Uint64ToString(value uint64) string {
	result := fmt.Sprintf("%d", value)
	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 整型64转为字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Int64ToString(value int64) string {
	result := fmt.Sprintf("%d", value)
	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 浮点64转为字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Float64ToString(value float64) string {
	result := fmt.Sprintf("%f", value)
	return result
}