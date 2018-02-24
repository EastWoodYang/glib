package glib

import (
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

/* ================================================================================
 * 随机数
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

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
