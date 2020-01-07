package glib

import (
	"fmt"
	"strings"
)

/* ================================================================================
 * 音乐
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 12音名转唱名
 * "C","#C","D","#D","E","F","#F","G","#G","A","#A","B"
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func MusicNamesToSingNames(musicNames []string) []string {
	singNames := make([]string, 0)
	nameTables := map[string]string{
		"C":  "1",
		"#C": "2",
		"bD": "2",
		"D":  "3",
		"#D": "4",
		"bE": "4",
		"E":  "5",
		"F":  "6",
		"#F": "7",
		"bG": "7",
		"G":  "8",
		"#G": "9",
		"bA": "9",
		"A":  "10",
		"#A": "11",
		"bB": "11",
		"B":  "12",
	}

	for _, musicName := range musicNames {
		if singName, isOk := nameTables[string(musicName)]; isOk {
			singNames = append(singNames, singName)
		}
	}

	return singNames
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取音名五线谱位置索引
 * 11
 * --------------------- 10
 * 9
 * --------------------- 8
 * 7
 * --------------------- 6
 * 5
 * --------------------- 4
 * 3
 * --------------------- 2
 * 1
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetMusicNameLocations(musicName string) []int {
	if len(musicName) == 1 {
		musicName = strings.ToUpper(musicName)
	} else if len(musicName) == 2 {
		musicName = fmt.Sprintf("%s%s", strings.ToLower(string(musicName[0])), strings.ToUpper(string(musicName[1])))
	}

	nameLocationTables := make(map[string][]int, 0)
	nameLocationTables["bC"] = []int{7}
	nameLocationTables["C"] = []int{7}
	nameLocationTables["#C"] = []int{7}
	nameLocationTables["bD"] = []int{1, 8}
	nameLocationTables["D"] = []int{1, 8}
	nameLocationTables["#D"] = []int{1, 8}
	nameLocationTables["bE"] = []int{2, 9}
	nameLocationTables["E"] = []int{2, 9}
	nameLocationTables["F"] = []int{3, 10}
	nameLocationTables["#F"] = []int{3, 10}
	nameLocationTables["bG"] = []int{4, 11}
	nameLocationTables["G"] = []int{4, 11}
	nameLocationTables["#G"] = []int{4, 11}
	nameLocationTables["bA"] = []int{5}
	nameLocationTables["A"] = []int{5}
	nameLocationTables["#A"] = []int{5}
	nameLocationTables["bB"] = []int{6}
	nameLocationTables["B"] = []int{6}

	locationIndexs := make([]int, 0)
	if locations, isOk := nameLocationTables[musicName]; isOk {
		locationIndexs = locations
	}

	return locationIndexs
}
