package glib

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符集合的交集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringInter(one, two []string) []string {
	allMap := make(map[string]bool, 0)
	interSet := make([]string, 0)

	if len(one) == 0 && len(two) == 0 {
		return interSet
	}

	if len(one) == 0 {
		for _, v := range two {
			allMap[v] = true
		}
	} else if len(two) == 0 {
		for _, v := range one {
			allMap[v] = true
		}
	} else {
		for _, v := range one {
			allMap[v] = true
		}

		for _, v := range two {
			if _, isOk := allMap[v]; isOk {
				allMap[v] = false
			}
		}
	}

	for k, v := range allMap {
		if !v {
			interSet = append(interSet, k)
		}
	}

	return interSet
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符集合的并集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringUnion(one, two []string) []string {
	allMap := make(map[string]string, 0)
	union := make([]string, 0)

	for _, v := range one {
		allMap[v] = v
	}

	for _, v := range two {
		if _, isOk := allMap[v]; !isOk {
			allMap[v] = v
		}
	}

	for _, v := range allMap {
		union = append(union, v)
	}

	return union
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 字符集合的差集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringDiff(one, two []string) []string {
	//并集合
	union := StringUnion(one, two)

	//交集合
	inter := StringInter(one, two)

	//差集合
	diff := FilterStringSlice(union, inter)

	return diff
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * int64交集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Int64Inter(one, two []int64) []int64 {
	allMap := make(map[int64]bool, 0)
	interSet := make([]int64, 0)

	if len(one) == 0 && len(two) == 0 {
		return interSet
	}

	if len(one) == 0 {
		for _, v := range two {
			allMap[v] = true
		}
	} else if len(two) == 0 {
		for _, v := range one {
			allMap[v] = true
		}
	} else {
		for _, v := range one {
			allMap[v] = true
		}
		for _, v := range two {
			if _, isOk := allMap[v]; isOk {
				allMap[v] = false
			}
		}
	}

	for k, v := range allMap {
		if !v {
			interSet = append(interSet, k)
		}
	}

	return interSet
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * int64并集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Int64Union(one, two []int64) []int64 {
	allMap := make(map[int64]int64, 0)
	union := make([]int64, 0)

	for _, v := range one {
		allMap[v] = v
	}

	for _, v := range two {
		if _, isOk := allMap[v]; !isOk {
			allMap[v] = v
		}
	}

	for _, v := range allMap {
		union = append(union, v)
	}

	return union
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * int64差集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Int64Diff(one, two []int64) []int64 {
	//并集合
	union := Int64Union(one, two)

	//交集合
	inter := Int64Inter(one, two)

	//差集合
	diff := FilterInt64Slice(union, inter)

	return diff
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * uint64交集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Uint64Inter(one, two []uint64) []uint64 {
	allMap := make(map[uint64]bool, 0)
	interSet := make([]uint64, 0)

	if len(one) == 0 && len(two) == 0 {
		return interSet
	}

	if len(one) == 0 {
		for _, v := range two {
			allMap[v] = true
		}
	} else if len(two) == 0 {
		for _, v := range one {
			allMap[v] = true
		}
	} else {
		for _, v := range one {
			allMap[v] = true
		}
		for _, v := range two {
			if _, isOk := allMap[v]; isOk {
				allMap[v] = false
			}
		}
	}

	for k, v := range allMap {
		if !v {
			interSet = append(interSet, k)
		}
	}

	return interSet
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * uint64并集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Uint64Union(one, two []uint64) []uint64 {
	allMap := make(map[uint64]uint64, 0)
	union := make([]uint64, 0)

	for _, v := range one {
		allMap[v] = v
	}

	for _, v := range two {
		if _, isOk := allMap[v]; !isOk {
			allMap[v] = v
		}
	}

	for _, v := range allMap {
		union = append(union, v)
	}

	return union
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * uint64差集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Uint64Diff(one, two []uint64) []uint64 {
	//并集合
	union := Uint64Union(one, two)

	//交集合
	inter := Uint64Inter(one, two)

	//差集合
	diff := FilterUint64Slice(union, inter)

	return diff
}




