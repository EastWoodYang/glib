package glib

import (
	"strings"
)

/* ================================================================================
 * String Table
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */
type (
	stringTable struct {
		tables []string
		flag   string
		count  int
	}
)

const (
	stringTables string = `ZBnLFW8MG7CEhXfS3sKgk4iV5ImJ6uDp1PjRt0ebqTrUv2wYzAxQdOyHa9Nc`
)

var (
	stringTableInstance *stringTable
)

func init() {
	stringTableInstance = NewStringTable()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化实例
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewStringTable(args ...int) *stringTable {
	count := 10
	if len(args) > 0 {
		count = args[0]

		if count <= 0 {
			count = 10
		}
	}

	return &stringTable{
		tables: strings.Split(stringTables, ""),
		flag:   "o",
		count:  count,
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数字映射字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *stringTable) NumberToStringTable(id int64) string {
	results := make([]string, 0, len(s.tables))
	length := int64(len(s.tables))
	index := 0

	for id/length > 0 {
		index = int(id % length)
		results = append(results, s.tables[index])

		id = id / length
	}

	index = int(id % length)
	results = append(results, s.tables[index])

	reverseStrings := ReverseString(strings.Join(results, ""))
	results = strings.Split(reverseStrings, "")

	if len(results) < s.count {
		results = append(results, s.flag)

		paddingLength := s.count - len(results)
		for i := 0; i < paddingLength; i++ {
			rndIndex := RandInt(int(length))
			results = append(results, s.tables[rndIndex])
		}
	}

	return strings.Join(results, "")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符串映射数字
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *stringTable) StringTableToNumber(sourceString string) int64 {
	var id int64
	length := int64(len(s.tables))
	codes := strings.Split(sourceString, "")

	for i, code := range codes {
		if code == s.flag {
			break
		}

		tableIndex := 0
		for j := 0; j < int(length); j++ {
			if code == s.tables[j] {
				tableIndex = j
				break
			}
		}

		if i > 0 {
			id = int64(id)*length + int64(tableIndex)
		} else {
			id = int64(tableIndex)
		}
	}

	return id
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数字映射字符串
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NumberToStringTable(id int64) string {
	return stringTableInstance.NumberToStringTable(id)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符串映射数字
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringTableToNumber(sourceString string) int64 {
	return stringTableInstance.StringTableToNumber(sourceString)
}
