package glib

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"net/url"
	"sort"
	"fmt"
	"strings"
	"strconv"
	"time"
)

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

	jsonDecoder := json.NewDecoder(bytes.NewBuffer(bytesData))
	jsonDecoder.UseNumber()

	return jsonDecoder.Decode(&object)
	//return json.Unmarshal(bytesData, object)
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