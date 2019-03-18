package glib

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

/* ================================================================================
 * 常用
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 捕获函数执行时的异常
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Try(fnSource func(), fnError func(interface{})) (er error) {
	defer func() {
		if err := recover(); err != nil {
			fnError(err)
			if err, ok := err.(error); ok {
				er = err
			} else {
				er = errors.New("try func call errors")
			}
		}
	}()

	fnSource()

	return er
}

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
func GetSubString(sourceString string, count int) string {
	if len(sourceString) <= count {
		return sourceString
	}

	newString, sourceStringRune := "", []rune(sourceString)
	sl, rl := 0, 0
	for _, r := range sourceStringRune {
		rint := int(r)
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}

		if sl+rl > count {
			break
		}
		sl += rl
		newString += string(r)
	}
	return newString
}

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
 * 翻转uint64 Slice
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func ReverseUint64Slice(int64Slice []uint64) {
	for from, to := 0, len(int64Slice)-1; from < to; from, to = from+1, to-1 {
		int64Slice[from], int64Slice[to] = int64Slice[to], int64Slice[from]
	}
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

	floatValue, err := strconv.ParseFloat(stringValue, 10)
	if err != nil {
		return 0.0
	}

	return floatValue
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
 * 过滤uint64数组（从all中过滤所有other中的数据，返回未被过滤的数据集合）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func FilterUint64Slice(all, other []uint64) []uint64 {
	allMap := make(map[uint64]bool, 0)
	diffSet := make([]uint64, 0)

	for _, a_v := range all {
		allMap[a_v] = true
	}

	for _, a_v := range other {
		if _, isOk := allMap[a_v]; isOk {
			allMap[a_v] = false
		}
	}

	for k, v := range allMap {
		if v {
			diffSet = append(diffSet, k)
		}
	}

	return diffSet
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 过滤字符数组（从all中过滤所有other中的数据，返回未被过滤的数据集合）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func FilterStringSlice(all, other []string) []string {
	allMaps := make(map[string]bool, 0)
	diffSet := make([]string, 0)

	for _, v := range all {
		allMaps[v] = true
	}

	for _, v := range other {
		if _, isOk := allMaps[v]; isOk {
			allMaps[v] = false
		}
	}

	for k, v := range allMaps {
		if v {
			diffSet = append(diffSet, k)
		}
	}

	return diffSet
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符集合的交集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringInter(one, two []string) []string {
	allMap := make(map[string]bool, 0)
	interSet := make([]string, 0)

	if len(one) == 0 && len(two) == 0 {
		return interSet
	}

	if len(one) == 0 {
		for _, v := range two {
			allMap[v] = true
		}
	} else if len(two) == 0 {
		for _, v := range one {
			allMap[v] = true
		}
	} else {
		for _, v := range one {
			allMap[v] = true
		}

		for _, v := range two {
			if _, isOk := allMap[v]; isOk {
				allMap[v] = false
			}
		}
	}

	for k, v := range allMap {
		if !v {
			interSet = append(interSet, k)
		}
	}

	return interSet
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符集合的并集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringUnion(one, two []string) []string {
	allMap := make(map[string]string, 0)
	union := make([]string, 0)

	for _, v := range one {
		allMap[v] = v
	}

	for _, v := range two {
		if _, isOk := allMap[v]; !isOk {
			allMap[v] = v
		}
	}

	for _, v := range allMap {
		union = append(union, v)
	}

	return union
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符集合的差集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringDiff(one, two []string) []string {
	//并集合
	union := StringUnion(one, two)

	//交集合
	inter := StringInter(one, two)

	//差集合
	diff := FilterStringSlice(union, inter)

	return diff
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * uint64交集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Uint64Inter(one, two []uint64) []uint64 {
	allMap := make(map[uint64]bool, 0)
	interSet := make([]uint64, 0)

	if len(one) == 0 && len(two) == 0 {
		return interSet
	}

	if len(one) == 0 {
		for _, v := range two {
			allMap[v] = true
		}
	} else if len(two) == 0 {
		for _, v := range one {
			allMap[v] = true
		}
	} else {
		for _, v := range one {
			allMap[v] = true
		}
		for _, v := range two {
			if _, isOk := allMap[v]; isOk {
				allMap[v] = false
			}
		}
	}

	for k, v := range allMap {
		if !v {
			interSet = append(interSet, k)
		}
	}

	return interSet
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * uint64并集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Uint64Union(one, two []uint64) []uint64 {
	allMap := make(map[uint64]uint64, 0)
	union := make([]uint64, 0)

	for _, v := range one {
		allMap[v] = v
	}

	for _, v := range two {
		if _, isOk := allMap[v]; !isOk {
			allMap[v] = v
		}
	}

	for _, v := range allMap {
		union = append(union, v)
	}

	return union
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * uint64差集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Uint64Diff(one, two []uint64) []uint64 {
	//并集合
	union := Uint64Union(one, two)

	//交集合
	inter := Uint64Inter(one, two)

	//差集合
	diff := FilterUint64Slice(union, inter)

	return diff
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 保留浮点数指定长度的小数位数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func ModFloat64(value float64, length int) float64 {
	format := fmt.Sprintf("0.%df", length)
	format = "%" + format
	valueString := fmt.Sprintf(format, value)
	floatValue, _ := strconv.ParseFloat(valueString, 64)
	return floatValue
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Base64字符编码
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func ToBase64(data string, args ...bool) string {
	resultString := ""
	isUrlEncoding := false

	if len(args) > 0 && args[0] {
		isUrlEncoding = true
	}

	if isUrlEncoding {
		resultString = base64.URLEncoding.EncodeToString([]byte(data))
	} else {
		resultString = base64.StdEncoding.EncodeToString([]byte(data))
	}

	return resultString
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Base64字符解码
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func FromBase64(data string, args ...bool) (string, error) {
	var bytesData []byte
	var err error
	isUrlDecoding := false

	if len(args) > 0 && args[0] {
		isUrlDecoding = true
	}

	if isUrlDecoding {
		bytesData, err = base64.URLEncoding.DecodeString(data)
	} else {
		bytesData, err = base64.StdEncoding.DecodeString(data)
	}

	if err != nil {
		return "", err
	}

	return string(bytesData), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 对象转换成Json字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func ToJson(object interface{}) (string, error) {
	v, err := json.Marshal(object)
	if err != nil {
		return "", err
	}

	return string(v), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Json字符串转换成对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func FromJson(jsonString string, object interface{}) error {
	bytesData := []byte(jsonString)
	return json.Unmarshal(bytesData, object)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 对象转换成Xml字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func ToXml(object interface{}) (string, error) {
	v, err := xml.Marshal(object)
	if err != nil {
		return "", err
	}

	return string(v), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Xml字符串转换成对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func FromXml(xmlString string, object interface{}) error {
	bytesData := []byte(xmlString)
	return xml.Unmarshal(bytesData, object)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字典参数升序排序，组成键＝值集合，然后把集合用&拼接成字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func JoinMapToString(params map[string]string, filterKeys []string, isEscape bool) string {
	var keys []string = make([]string, 0)
	var values []string = make([]string, 0)

	filterKeyMaps := make(map[string]string, 0)
	if len(filterKeys) > 0 {
		for _, key := range filterKeys {
			filterKeyMaps[key] = key
		}
	}

	//请求参数排序（字母升序）
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	//拼接KeyValue字符串
	for _, key := range keys {
		//过滤空值
		if len(params[key]) > 0 {
			//过滤指定的key
			if _, isExists := filterKeyMaps[key]; isExists {
				continue
			}

			keyValue := params[key]
			if isEscape {
				keyValue = url.QueryEscape(keyValue)
			}

			//键＝值集合
			value := fmt.Sprintf("%s=%s", key, keyValue)
			values = append(values, value)
		}
	}

	//用&连接起来
	paramString := strings.Join(values, "&")

	return paramString
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 签名算法
 * params里的每个Value都需要进行url编码
 * fmt.Sprintf("%s=%s", key, url.QueryEscape(value))
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func MapDataSign(params map[string]string, secret string) (string, bool) {
	isExpired := false

	var timestamp int64
	var keys []string = make([]string, 0)
	var values []string = make([]string, 0)

	//请求参数排序（字母升序）
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	//拼接KeyValue字符串
	for _, key := range keys {
		if len(params[key]) > 0 {
			values = append(values, key)         //Key
			values = append(values, params[key]) //Value

			if key == "timestamp" {
				if _timestamp, err := strconv.ParseInt(params[key], 10, 64); err != nil {
					timestamp = _timestamp
				}
			}
		}
	}
	paramString := strings.Join(values, "")

	//是否已过期
	if timestamp > 0 {
		isExpired = time.Unix(timestamp, 0).Add(time.Minute * time.Duration(5)).Before(time.Now())
	}

	//Md5签名（在拼接的字符串头尾附加上api密匙，然后md5，md5串是大写）
	paramString = fmt.Sprintf("%s%s%s", secret, paramString, secret)
	sign := Md5(paramString)

	return strings.ToUpper(sign), isExpired
}
