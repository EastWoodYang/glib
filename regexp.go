package glib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/* ================================================================================
 * 正则表达式
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊
 * ================================================================================ */

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
	pattern := "^[\u4E00-\u9FFF\uFF00-\uFFEF]$"
	if len(args) == 1 {
		pattern = fmt.Sprintf("^[\u4E00-\u9FFF\uFF00-\uFFEF]{1,%d}$", args...)
	} else if len(args) == 2 {
		pattern = fmt.Sprintf("^[\u4E00-\u9FFF\uFF00-\uFFEF]{%d,%d}$", args...)
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
		pattern = fmt.Sprintf("^%s{1,%d}$", items...)
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
		pattern = fmt.Sprintf("^\\w{1,%d}$", args...)
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
		pattern = fmt.Sprintf("^\\d{1,%d}$", args...)
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
	pattern := "^0?(\\d{2})?1[3|4|5|6|7|8|9][0-9]\\d{8}$"
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
 * 判断是否sql
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func IsSql(source string) bool {
	if source == "" {
		return false
	}

	pattern := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, _ := regexp.Compile(pattern)

	return re.MatchString(source)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * sql过滤
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func SqlFilter(source string) string {
	if source == "" {
		return source
	}

	pattern := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, _ := regexp.Compile(pattern)

	return re.ReplaceAllString(source, "")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 去除连续的换行符
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func TrimSpaceLine(source string) string {
	if source == "" {
		return source
	}

	re, _ := regexp.Compile("\\s{2,}")
	trimFunc := re.ReplaceAllString(source, "\r\n")

	return strings.TrimSpace(trimFunc)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 换行转br标签
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func String2Br(source string) string {
	if source == "" {
		return source
	}

	re, _ := regexp.Compile("\r\n|\n")

	return re.ReplaceAllString(source, "<br />")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 换行转br标签
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Br2String(source string) string {
	if source == "" {
		return source
	}

	re, _ := regexp.Compile("<br />")

	return re.ReplaceAllString(source, "\r\n")
}
