package glib

/* ================================================================================
 * 音乐
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 音名转唱名
 * "C","#C","D","#D","E","F","#F","G","#G","A","#A","B"
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func MusicNamesToSingNames(musicCodes []string) []string {
	singCodes := make([]string, 0)
	musicTables := map[string]string{
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

	for _, music := range musicCodes {
		if sing, isOk := musicTables[string(music)]; isOk {
			singCodes = append(singCodes, sing)
		}
	}

	return singCodes
}
