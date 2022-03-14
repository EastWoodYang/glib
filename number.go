package glib

import (
	"fmt"
	"strconv"
)

/* ================================================================================
 * 数字
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

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