package glib

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/* ================================================================================
 * 日期
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断firstDatetime是否在secondDatetime的后面，即firstDatetime比secondDatetime日期大
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func IsDateGreaterThan(firstDatetime, secondDatetime time.Time) bool {
	return firstDatetime.After(secondDatetime)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断firstDatetime是否在secondDatetime的前面，即firstDatetime比secondDatetime日期小
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func IsDateLessThan(firstDatetime, secondDatetime time.Time) bool {
	return firstDatetime.Before(secondDatetime)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前Unix秒时间戳
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func UnixTimestamp(args ...bool) int64 {
	return GetNow(args...).Unix()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前Unix纳秒时间戳
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func UnixNanoTimestamp(args ...bool) int64 {
	return GetNow(args...).UnixNano()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Unix日期（1970-01-01 00:00:00）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func UnixTimestampDate() time.Time {
	dtTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local)
	return dtTime
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Unix日期（0001-01-01 00:00:00）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func UnixDate() time.Time {
	dtTime, _ := StringToTime(time.UnixDate)
	return dtTime
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据Unix时间戳返回日期
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func UnixTimestampToDate(unixTimestamp int64) time.Time {
	return time.Unix(unixTimestamp, 0)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据Unix纳秒时间戳返回日期
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func UnixNanoTimestampToDate(unixNanoTimestamp int64) time.Time {
	return time.Unix(0, unixNanoTimestamp)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前日期时间
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetNow(args ...bool) time.Time {
	isUtc := false
	if len(args) > 0 {
		isUtc = args[0]
	}
	timeNow := time.Now()
	if isUtc {
		timeNow = time.Now().UTC()
	}

	return timeNow
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前年份
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetCurrentYear(args ...bool) int32 {
	year, _, _ := GetNow(args...).Date()
	return int32(year)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前月份
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetCurrentMonth(args ...bool) int32 {
	_, month, _ := GetNow(args...).Date()
	return int32(month)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前日
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetCurrentDay(args ...bool) int32 {
	_, _, day := GetNow(args...).Date()
	return int32(day)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前小时
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetCurrentHour(args ...bool) int32 {
	hour, _, _ := GetNow(args...).Clock()
	return int32(hour)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前分钟
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetCurrentMinute(args ...bool) int32 {
	_, minute, _ := GetNow(args...).Clock()
	return int32(minute)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取当前秒数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetCurrentSecond(args ...bool) int32 {
	_, _, second := GetNow(args...).Clock()
	return int32(second)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取日期时间的年份
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetDateYear(datetime time.Time) int32 {
	year, _, _ := datetime.Date()
	return int32(year)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取日期时间的月份
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetDateMonth(datetime time.Time) int32 {
	_, month, _ := datetime.Date()
	return int32(month)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取日期时间的日部分
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetDateDay(datetime time.Time) int32 {
	_, _, day := datetime.Date()
	return int32(day)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取日期时间的小时部分
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetDateHour(datetime time.Time) int32 {
	hour, _, _ := datetime.Clock()
	return int32(hour)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取日期时间的分钟部分
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetDateMinute(datetime time.Time) int32 {
	_, minute, _ := datetime.Clock()
	return int32(minute)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取日期时间的秒部分
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetDateSecond(datetime time.Time) int32 {
	_, _, second := datetime.Clock()
	return int32(second)
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
 * 获取日期时间的日期和星期字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetDatetimeWeekString(datetime time.Time) string {
	_, month, day := datetime.Date()
	hour, minute, _ := datetime.Clock()
	weekday := GetWeek(datetime)

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
 * 获取友好的日期时间字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func TimeToFriendString(datetime time.Time, args ...string) string {
	format := "2006-01-02 15:04:05"
	if len(args) > 0 {
		format = args[0]
	}

	result := TimeToString(datetime, format)
	currentDate := time.Now()
	year1, month1, day1 := currentDate.Date()
	year2, month2, day2 := datetime.Date()

	dayCount := 0
	if year1 == year2 {
		if month1 == month2 {
			dayCount = day1 - day2
			if dayCount == 0 {
				hour := currentDate.Hour() - datetime.Hour()
				if hour == 0 {
					minutesCount := currentDate.Minute() - datetime.Minute()
					if minutesCount == 0 {
						result = "刚刚"
					} else {
						result = fmt.Sprintf("%d分钟前", minutesCount)
					}
				} else {
					result = fmt.Sprintf("%d小时前", hour)
				}
			} else {
				if dayCount > 14 {
					result = "半个月前"
				} else if dayCount > 6 {
					result = "一周前"
				} else {
					result = fmt.Sprintf("%d天前", dayCount)
				}
			}
		} else {
			dayCount = int(currentDate.Sub(datetime).Seconds() / 86400)
			if dayCount >= 90 && dayCount < 120 {
				result = "3个月前"
			} else if dayCount >= 60 && dayCount < 90 {
				result = "2个月前"
			} else if dayCount >= 30 && dayCount < 60 {
				result = "1个月前"
			}
		}
	}

	return result
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
 * firstDatetime加上时间间隔duration，返回日期时间
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func DatetimeAdd(firstDatetime time.Time, duration time.Duration) time.Time {
	return firstDatetime.Add(duration)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * firstDatetime减去secondDatetime，返回时间间隔
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func DatetimeSub(firstDatetime, secondDatetime time.Time) time.Duration {
	return firstDatetime.Sub(secondDatetime)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 在当前的日期时间增加指定的分钟数，返回日期时间
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func AddMinutesForCurrent(minutes int) time.Time {
	return time.Now().Add(time.Duration(minutes) * time.Minute)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 在指定的日期时间增加指定的分钟数，返回日期时间
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func AddMinutesForDatetime(datetime time.Time, minutes int) time.Time {
	return datetime.Add(time.Duration(minutes) * time.Minute)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 时间字符串加指定的分钟数，返回时间字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func TimeStringAddMinutes(timeString string, minutes int) string {
	format := "15:04:05"

	var timeValue time.Time
	if time, err := time.ParseInLocation(format, timeString, time.Local); err == nil {
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
	if time, err := time.ParseInLocation(format, timeString, time.Local); err == nil {
		timeValue = time
	}

	y, m, d := date.Date()
	h1, m1, s1 := timeValue.Clock()
	datetime := time.Date(y, m, d, h1, m1, s1, 0, time.Local)
	return datetime
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
 * 获取当前日期是周几（1:周一｜2:周二｜...|7:周日）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetCurrentWeek() int {
	return GetWeek(time.Now())
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取指定的日期是周几（1:周一｜2:周二｜...|7:周日）
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetWeek(date time.Time) int {
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
 * 获取当前周对应的月份里的日期范围（minDay in month, maxDay in month）
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
 * 获取日期范围内的所属周几的日期集合
 * week：从1开始，1表示周一，依次类推
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetDateRangeForWeekInDateRange(startDate, endDate time.Time, week int) []time.Time {
	dateList := make([]time.Time, 0)
	date := startDate

	for date.Before(endDate) || date.Equal(endDate) {
		weekValue := GetWeek(date)
		if weekValue == week {
			dateList = append(dateList, date)
		}

		date = date.AddDate(0, 0, 1)
	}

	return dateList
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
func CurrentTimeToString(args ...interface{}) string {
	return TimeToString(time.Now(), args...)
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
