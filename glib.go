package glib

import (
	"bufio"
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/md5"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

/* ================================================================================
 * 常用帮助模块
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊
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
 * 获取当前唯一Token字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetTokenString() string {
	timestamp := GetUnixTimestamp()
	return Md5(strconv.FormatInt(timestamp, 10))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前Unix时间戳
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetUnixTimestamp() int64 {
	return time.Now().Unix()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取字符串个数（不是字节数）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetStringCount(sourceString string) int {
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
 * 用指定的字符串把uint64切片链接为字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Uint64SliceToString(uintSlice []uint64, args ...string) string {
	result := ""
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
 * uint64差集
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetUint64DiffSet(all, other []uint64) []uint64 {
	allMaps := make(map[uint64]bool, 0)
	diffSet := make([]uint64, 0)

	for _, a_v := range all {
		allMaps[a_v] = true
	}

	for _, a_v := range other {
		if _, isOk := allMaps[a_v]; isOk {
			allMaps[a_v] = false
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
 * string差集
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetStringDiffSet(all, other []string) []string {
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
 * 获取一段时间范围内指定间隔的时间段集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetTimeIntervalStringSlice(startDate, endDate time.Time, minutes int64) []string {
	timeStringList := make([]string, 0)

	date := startDate
	for date.Before(endDate) || date.Equal(endDate) {
		timeString := TimeToString(date, "15:04")
		timeStringList = append(timeStringList, timeString)

		date = date.Add(time.Duration(minutes) * time.Minute)
	}

	return timeStringList
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取日期时间的日期和星期字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetDatetimeWeekString(datetime time.Time) string {
	_, month, day := datetime.Date()
	hour, minute, _ := datetime.Clock()
	weekday := GetWeekWithDate(datetime)

	weekdays := make(map[int]string, 0)
	weekdays[1] = "星期一"
	weekdays[2] = "星期二"
	weekdays[3] = "星期三"
	weekdays[4] = "星期四"
	weekdays[5] = "星期五"
	weekdays[6] = "星期六"
	weekdays[7] = "星期日"
	weekdayString := weekdays[weekday]

	dateString := fmt.Sprintf("%d月%d日%s%d:%d", int(month), day, weekdayString, hour, minute)

	return dateString
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Unix日期（1970-01-01 00:00:00）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func UnixDate() time.Time {
	dtTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local)
	//dtTime, _ := StringToTime(time.UnixDate)
	return dtTime
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 日期字符串切片转成日期
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringSliceToDate(dateStringSlice []string) (time.Time, error) {
	dateString := strings.Join(dateStringSlice, "-")

	dtTime, err := StringToTime(dateString)
	if err != nil {
		return UnixDate(), err
	}

	return GetMinDate(dtTime), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 日期转成日期字符串切片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func DateToStringSlice(date time.Time) []string {
	if date.IsZero() {
		date = time.Now()
	}

	dateStringSlice := make([]string, 0)
	year, month, day := date.Date()

	dateStringSlice = append(dateStringSlice, fmt.Sprintf("%d", year))
	dateStringSlice = append(dateStringSlice, fmt.Sprintf("%d", int(month)))
	dateStringSlice = append(dateStringSlice, fmt.Sprintf("%d", day))

	return dateStringSlice
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * int切片转成日期
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func IntSliceToDate(intSlice []int) (time.Time, error) {
	dateStringSlice := make([]string, 3)
	for _, v := range intSlice {
		dateStringSlice = append(dateStringSlice, fmt.Sprintf("%d", v))
	}

	dtTime, err := StringToTime(strings.Join(dateStringSlice, "-"))
	if err != nil {
		return UnixDate(), err
	}

	return GetMinDate(dtTime), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 日期转成int切片
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func DateToIntSlice(date time.Time) []int {
	intSlice := make([]int, 3)

	dateStringSlice := DateToStringSlice(date)
	intSlice[0], _ = strconv.Atoi(dateStringSlice[0])
	intSlice[1], _ = strconv.Atoi(dateStringSlice[1])
	intSlice[2], _ = strconv.Atoi(dateStringSlice[2])

	return intSlice
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 返回日期的最小日期时间（2016-01-02 00:00:00）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetMinDate(dtTime time.Time) time.Time {
	year, month, day := dtTime.Date()
	return time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.Local)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 返回日期的最大日期时间（2016-01-02 59:59:59 999）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetMaxDate(dtTime time.Time) time.Time {
	year, month, day := dtTime.Date()
	return time.Date(int(year), time.Month(month), int(day), 59, 59, 59, 999, time.Local)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 日期时间增加指定的分钟数，返回日期时间
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func DatetimeAddMinutes(datetime time.Time, minutes int) time.Time {
	return datetime.Add(time.Duration(minutes) * time.Minute)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 时间字符串加指定的分钟数，返回时间字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func TimeStringAddMinutes(timeString string, minutes int) string {
	format := "15:04:05"

	var timeValue time.Time
	if time, err := time.Parse(format, timeString); err == nil {
		timeValue = time
	}

	timeValue = timeValue.Add(time.Duration(minutes) * time.Minute)
	return timeValue.Format(format)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 日期时间的日期部分和时间字符串连接，返回日期时间
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetDatetimeForDateAndTimeString(date time.Time, timeString string) time.Time {
	format := "15:04:05"

	var timeValue time.Time
	if time, err := time.Parse(format, timeString); err == nil {
		timeValue = time
	}

	y, m, d := date.Date()
	h1, m1, s1 := timeValue.Clock()
	datetime := time.Date(y, m, d, h1, m1, s1, 0, time.Local)
	return datetime
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取日期范围内的所属周几的日期集合
 * week：从1开始，1表示周一，依次类推
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetDateRangeForWeekInDateRange(startDate, endDate time.Time, week int) []time.Time {
	dateList := make([]time.Time, 0)
	date := startDate

	for date.Before(endDate) || date.Equal(endDate) {
		weekValue := GetWeekWithDate(date)
		if weekValue == week {
			dateList = append(dateList, date)
		}

		date = date.AddDate(0, 0, 1)
	}

	return dateList
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前日期是周几（1:周一｜2:周二｜...|7:周日）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetCurrentWeek() int {
	nowDate := time.Now()
	return GetWeekWithDate(nowDate)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取指定的日期是周几（1:周一｜2:周二｜...|7:周日）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetWeekWithDate(date time.Time) int {
	nowDate := date
	days := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
		0: 7,
	}
	weekday := nowDate.Weekday() //0：周日 | 1：周一 | .. ｜6：周六
	weekdayValue := days[int(weekday)]

	return weekdayValue
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前周对应的日期范围（minDay in month, maxDay in month）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetCurrentWeekDayRange() (int, int) {
	nowDate := time.Now()
	_, _, day := nowDate.Date()
	weekdayValue := GetCurrentWeek()
	minDay := day - weekdayValue + 1
	maxDay := day + 7 - weekdayValue

	return minDay, maxDay
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前日期月份对应的天数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetCurrentDayCount() int {
	return GetDayCount(time.Now())
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取指定日期月份对应的天数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetDayCount(datetime time.Time) int {
	year, month, _ := datetime.Date()
	dayCount := 31
	if month == 4 || month == 6 || month == 9 || month == 11 {
		dayCount = 30
	} else if month == 2 {
		if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
			dayCount = 29
		} else {
			dayCount = 28
		}
	}

	return dayCount
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
 * 分钟数转时间字符串（HH:mm:ss）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func MinutesToTimeString(minutes int64) string {
	hoursPart := minutes / 60
	minutesPart := minutes % 60

	timeString := fmt.Sprintf("%02d:%02d:00", hoursPart, minutesPart)

	return timeString
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 时间转字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func TimeToString(timeValue time.Time, args ...interface{}) string {
	format := "2006-01-02 15:04:05"
	if len(args) == 1 {
		if v, ok := args[0].(string); ok {
			format = v
		}
	}
	timeStrng := timeValue.Format(format)
	return timeStrng
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符串转时间
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringToTime(timeString string, args ...interface{}) (time.Time, error) {
	format := "2006-01-02 15:04:05"
	if len(args) == 1 {
		if v, ok := args[0].(string); ok {
			format = v
		}
	}

	return time.ParseInLocation(format, timeString, time.Local)
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
 * 正则是否匹配
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func IsRegexpMatch(sourceString string, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(sourceString)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否中文
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func IsChinese(sourceString string, args ...interface{}) bool {
	pattern := "^[\u4E00-\u9FFF]$"
	if len(args) == 1 {
		pattern = fmt.Sprintf("^[\u4E00-\u9FFF]{,%d}$", args...)
	} else if len(args) == 2 {
		pattern = fmt.Sprintf("^[\u4E00-\u9FFF]{%d,%d}$", args...)
	}

	re := regexp.MustCompile(pattern)
	return re.MatchString(sourceString)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否身份证号码
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func IsIdCardNum(sourceString string) bool {
	pattern := "^\\d{15}$|^\\d{18}$|^\\d{17}(\\d|X|x)$"
	re := regexp.MustCompile(pattern)
	return re.MatchString(sourceString)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否用户名（以英文字母开头，后面跟英文字母和数据以及下划线）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func IsUsername(sourceString string, args ...interface{}) bool {
	pattern := "[a-zA-Z]{1}[a-zA-Z0-9_]"
	items := []interface{}{pattern}

	for _, item := range args {
		items = append(items, item)
	}

	if len(args) == 1 {
		pattern = fmt.Sprintf("^%s{,%d}$", items...)
	} else if len(args) == 2 {
		pattern = fmt.Sprintf("^%s{%d,%d}$", items...)
	} else {
		pattern = fmt.Sprintf("^%s{%d,%d}$", pattern, 5, 15)
	}
	re := regexp.MustCompile(pattern)
	return re.MatchString(sourceString)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否英文单词
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func IsAlpha(sourceString string, args ...interface{}) bool {
	pattern := "^\\w+$"
	if len(args) == 1 {
		pattern = fmt.Sprintf("^\\w{,%d}$", args...)
	} else if len(args) == 2 {
		pattern = fmt.Sprintf("^\\w{%d,%d}$", args...)
	}

	re := regexp.MustCompile(pattern)
	return re.MatchString(sourceString)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否数字
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func IsNumber(sourceString string, args ...interface{}) bool {
	pattern := "^\\d+$"
	if len(args) == 1 {
		pattern = fmt.Sprintf("^\\d{,%d}$", args...)
	} else if len(args) == 2 {
		pattern = fmt.Sprintf("^\\d{%d,%d}$", args...)
	}

	re := regexp.MustCompile(pattern)
	return re.MatchString(sourceString)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否英文单词或数字
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func IsAlphaOrNumber(sourceString string) bool {
	return IsAlpha(sourceString) || IsNumber(sourceString)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否电子邮件
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func IsEmail(sourceString string, args ...interface{}) bool {
	pattern := "^[a-zA-Z0-9]{1}[a-zA-Z0-9_-]*@[a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,}(\\.[a-zA-Z]+)+$"

	if len(args) == 1 {
		if length, err := strconv.Atoi(fmt.Sprintf("%d", args[0])); err != nil || len(sourceString) > length {
			return false
		}
	} else if len(args) == 2 {
		if minLength, err := strconv.Atoi(fmt.Sprintf("%d", args[0])); err != nil {
			return false
		} else if maxLength, err := strconv.Atoi(fmt.Sprintf("%d", args[1])); err != nil {
			return false
		} else {
			if length := len(sourceString); length < minLength || length > maxLength {
				return false
			}
		}
	}

	re := regexp.MustCompile(pattern)
	return re.MatchString(sourceString)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否手机号码
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func IsMobile(sourceString string) bool {
	//pattern := "^0?(\\d{2})?1[3|4|5|6|7|8][0-9]\\d{8}$"
	pattern := "^0?(\\d{2})?1[3|4|5|6|7|8|9][0-9]\\d{8}$" //apple 测试帐号199号段
	re := regexp.MustCompile(pattern)
	return re.MatchString(sourceString)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否电话号码
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func IsTelphone(sourceString string) bool {
	pattern := "^0\\d{2,3}-?\\d{7,8}$|^\\d{7,8}-?\\d{3,5}$"
	re := regexp.MustCompile(pattern)
	return re.MatchString(sourceString)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 唯一Guid
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Guid() string {
	data := make([]byte, 48)

	if _, err := io.ReadFull(crand.Reader, data); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(data))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 随机Salt
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Salt() string {
	return fmt.Sprintf("%d",
		RandInt(999999999),
	)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * maxInt以内的随机数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func RandInt(maxInt int) int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rnd.Intn(maxInt)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * min,max范围内的随机数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func RandIntRange(min, max int) int {
	return min + RandInt(max-min)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 整型数组随机打乱
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Shuffle(arr []int, args ...int) {
	n := 0
	argsCount := len(args)
	if argsCount == 0 {
		n = len(arr) - 1
	} else if argsCount == 1 {
		n = args[0]
	}

	if n <= 0 {
		return
	}

	Shuffle(arr, n-1)
	rndIndex := RandIntRange(0, n)

	arr[n], arr[rndIndex] = arr[rndIndex], arr[n]
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 随机英文和数字字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func RandString(length int) string {
	rawStrings := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "")
	results := make([]string, length)

	for i := 0; i < length; i++ {
		index := RandInt(len(rawStrings) - 1)
		results[i] = rawStrings[index]
	}

	return strings.Join(results, "")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 指定长度的随机数字
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func RandNumberString(length int) string {
	rawStrings := strings.Split("0123456789", "")
	results := make([]string, length)

	for i := 0; i < length; i++ {
		index := RandInt(len(rawStrings) - 1)
		results[i] = rawStrings[index]
	}

	return strings.Join(results, "")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取url里指定参数的值
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetUrlParamValue(url, paramName string) string {
	paramValue := ""
	urlParams := strings.Split(url, "?")

	if len(urlParams) == 2 {
		params := strings.Split(urlParams[1], "&")
		for _, v := range params {
			items := strings.Split(v, "=")
			if strings.ToLower(items[0]) == strings.ToLower(paramName) {
				paramValue = items[1]
				break
			}
		}
	}

	return paramValue
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取数据的缓存key
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetModelKey(model interface{}, prefixKey, fieldName string) string {
	typeOf := reflect.TypeOf(model)
	valueOf := reflect.ValueOf(model)
	valueElem := valueOf.Elem()

	if len(fieldName) == 0 {
		fieldName = "Id"
	}

	if kind := typeOf.Kind(); kind != reflect.Ptr {
		panic("Model is not a pointer type")
	}

	modelKey := ""
	pkgName := strings.Split(valueElem.String(), " ")[0][1:]

	if _, ok := valueElem.Type().FieldByName(fieldName); !ok {
		panic("Model does not contain Id field")
	}

	fieldValue := valueElem.FieldByName(fieldName).Uint()
	modelKey = fmt.Sprintf("%s||%s||%d", prefixKey, pkgName, fieldValue)
	modelKey = strings.ToLower(modelKey)

	return modelKey
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Md5哈希
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Md5(data string) string {
	m := md5.New()
	io.WriteString(m, data)
	return hex.EncodeToString(m.Sum(nil))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Sha1哈希
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Hmac Sha1哈希
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HmacSha1(data string, key string, args ...bool) string {
	resultString := ""
	isHex := true

	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))

	if len(args) > 0 {
		isHex = args[0]
	}

	if isHex {
		//resultString = fmt.Sprintf("%x", mac.Sum(nil))
		resultString = hex.EncodeToString(mac.Sum(nil))
	} else {
		resultString = string(mac.Sum(nil))
	}
	return resultString
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Sha1哈希
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Sha256(data string) string {
	t := sha256.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Sha256WithRsa签名
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Sha256WithRsa(data string, privateKey string) (string, error) {
	//key to pem
	privateKey = RsaPrivateToMultipleLine(privateKey)

	//RSA密匙
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("sign private key decode error")
	}

	prk8, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	//SHA256哈希
	h := sha256.New()
	h.Write([]byte(data))
	digest := h.Sum(nil)

	//RSA签名
	rsaPrivateKey := prk8.(*rsa.PrivateKey)
	s, err := rsa.SignPKCS1v15(nil, rsaPrivateKey, crypto.SHA256, []byte(digest))
	if err != nil {
		return "", err
	}
	sign := base64.StdEncoding.EncodeToString(s)
	return string(sign), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Sha256WithRsa验签
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Sha256WithRsaVerify(data, sign, publicKey string) (bool, error) {
	// publickey to pem
	publicKey = RsaPublicToMultipleLine(publicKey)

	// 加载RSA的公钥
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return false, errors.New("sign public key decode error")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}
	rsaPublicKey, _ := pub.(*rsa.PublicKey)

	h := sha256.New()
	h.Write([]byte(data))
	digest := h.Sum(nil)

	// base64 decode,支付宝对返回的签名做过base64 encode必须要反过来decode才能通过验证
	signString, _ := base64.StdEncoding.DecodeString(sign)
	hex.EncodeToString(signString)

	// 调用rsa包的VerifyPKCS1v15验证签名有效性
	if err = rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, digest, signString); err != nil {
		return false, err
	}

	return true, nil
}

/*
func testAes() {
	key := []byte("axewfd3_r44&98Klaxewfd3_r44&98Kl")
	result, err := AESEncrypt([]byte("mliu"), key)
	if err != nil {
		panic(err)
	}
	s := base64.StdEncoding.EncodeToString(result)
	origData, err := AESDecrypt(result, key)
	ss := string(origData)
}*/

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Des加密
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key[0:8])
	if err != nil {
		return nil, err
	}
	origData = Pkcs5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Des解密
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key[0:8])
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = Pkcs5UnPadding(origData)
	return origData, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Aes加密
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key[0:16])
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = Pkcs5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Aes解密
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key[0:16])
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))

	blockMode.CryptBlocks(origData, crypted)
	origData = Pkcs5UnPadding(origData)
	return origData, nil
}

func Pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func Pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Rsa加密
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func RSAEncrypt(origData, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(crand.Reader, pub, origData)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Rsa解密
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func RSADecrypt(ciphertext, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(crand.Reader, priv, ciphertext)
}

func GenRsaKey(bits int) (string, string, error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(crand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	privateBuffer := bytes.NewBuffer(make([]byte, 0))
	privateWriter := bufio.NewWriter(privateBuffer)
	err = pem.Encode(privateWriter, block)
	if err != nil {
		return "", "", err
	}

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	publicBuffer := bytes.NewBuffer(make([]byte, 0))
	publicWriter := bufio.NewWriter(publicBuffer)
	err = pem.Encode(publicWriter, block)
	if err != nil {
		return "", "", err
	}
	return privateBuffer.String(), publicBuffer.String(), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 将单行的Ras Public字符串格式化为多行格式
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func RsaPublicToMultipleLine(privateKey string) string {
	privateKeys := make([]string, 0)
	privateKeys = append(privateKeys, "-----BEGIN PUBLIC KEY-----")

	for i := 0; i < 4; i++ {
		start := i * 64
		end := (i + 1) * 64
		lineKey := ""
		if i == 3 {
			lineKey = privateKey[start:]
		} else {
			lineKey = privateKey[start:end]
		}
		privateKeys = append(privateKeys, lineKey)
	}

	privateKeys = append(privateKeys, "-----END PUBLIC KEY-----")
	privateKey = strings.Join(privateKeys, "\r\n")

	return privateKey
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 将单行的Rsa Private字符串格式化为多行格式
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func RsaPrivateToMultipleLine(privateKey string) string {
	privateKeys := make([]string, 0)
	privateKeys = append(privateKeys, "-----BEGIN PRIVATE KEY-----")

	for i := 0; i < 26; i++ {
		start := i * 64
		end := (i + 1) * 64
		lineKey := ""
		if i == 25 {
			lineKey = privateKey[start:]
		} else {
			lineKey = privateKey[start:end]
		}
		privateKeys = append(privateKeys, lineKey)
	}

	privateKeys = append(privateKeys, "-----END PRIVATE KEY-----")
	privateKey = strings.Join(privateKeys, "\r\n")

	return privateKey
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
