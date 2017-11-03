package glib

import (
	"regexp"
	"strings"
)

/* ================================================================================
 * 常用
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 安全字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func SafeString(source string, args ...bool) string {
	if source == "" {
		return source
	}

	isFilterHtmlTag := true
	isFilterSql := true
	argCount := len(args)
	if argCount > 0 {
		if argCount == 1 {
			isFilterHtmlTag = args[0]
		}
		if argCount > 1 {
			isFilterSql = args[1]
		}
	}

	content := ""
	if isFilterHtmlTag || isFilterSql {
		if isFilterHtmlTag {
			content = HtmlTagFilter(source)
		}

		if isFilterSql {
			content = SqlFilter(source)
		}
	}

	content = HtmlEncode(source)

	return content
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 安全参数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func SafeParam(source string) string {
	if source == "" {
		return source
	}

	content := SqlFilter(source)
	return content
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * html编码
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HtmlEncode(source string) string {
	result := source
	if result == "" {
		return ""
	}

	all := make(map[string]string, 0)
	all[">"] = "&lt;"
	all["<"] = "&gt;"
	all["&"] = "&amp;"
	all["\""] = "&quot;"
	all["'"] = "&#39;"
	//all[" "] = "&nbsp;"

	for k, v := range all {
		result = strings.Replace(result, k, v, -1)
	}

	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * html解码
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HtmlDecode(source string) string {
	result := source
	if result == "" {
		return ""
	}

	all := make(map[string]string, 0)
	all[">"] = "&lt;"
	all["<"] = "&gt;"
	all["&"] = "&amp;"
	all["\""] = "&quot;"
	all["'"] = "&#39;"
	//all[" "] = "&nbsp;"

	for k, v := range all {
		result = strings.Replace(result, v, k, -1)
	}

	for k, v := range all {
		result = strings.Replace(result, v, k, -1)
	}

	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 过滤html标签
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HtmlTagFilter(source string) string {
	if source == "" {
		return source
	}

	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")

	return re.ReplaceAllString(source, "")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 过滤html 链接标签
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HtmlHyperLinkFilter(source string) string {
	if source == "" {
		return source
	}

	//re, _ := regexp.Compile("\\<a[\\s][\\S\\s]+?/\\>|\\<a[\\S\\s]+?\\</a\\>")
	re, _ := regexp.Compile("\\<a[\\s][\\S\\s]+?/\\>|\\<a[\\s][\\S\\s]+?\\</a\\>")

	return re.ReplaceAllString(source, "")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 过滤html img标签
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HtmlImageFilter(source string) string {
	if source == "" {
		return source
	}

	re, _ := regexp.Compile("\\<img[\\S\\s]+?/\\>")

	return re.ReplaceAllString(source, "")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 过滤html audio标签
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HtmlAudioFilter(source string) string {
	if source == "" {
		return source
	}

	re, _ := regexp.Compile("\\<audio[\\S\\s]+?/\\>|\\<audio[\\S\\s]+?\\</audio\\>")

	return re.ReplaceAllString(source, "")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 过滤html video标签
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HtmlVideoFilter(source string) string {
	if source == "" {
		return source
	}

	re, _ := regexp.Compile("\\<video[\\S\\s]+?/\\>|\\<video[\\S\\s]+?\\</video\\>")

	return re.ReplaceAllString(source, "")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 过滤html css标签
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HtmlCssFilter(source string) string {
	if source == "" {
		return source
	}

	re, _ := regexp.Compile("\\<style[\\S\\s]+?/\\>|\\<style[\\S\\s]+?\\</style\\>")

	return re.ReplaceAllString(source, "")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 过滤html script标签
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HtmlScriptFilter(source string) string {
	if source == "" {
		return source
	}

	//\\<!--[^>]+\\>
	//\\<script[\\S\\s]+?/\\>

	re, _ := regexp.Compile("\\<!--[^>]+\\>|\\<iframe[\\S\\s]+?/\\>|\\<iframe[\\S\\s]+?\\</iframe\\>|\\<script[\\S\\s]+?/\\>|\\<script[\\S\\s]+?\\</script\\>")

	return re.ReplaceAllString(source, "")
}
