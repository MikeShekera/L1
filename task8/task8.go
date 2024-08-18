package task8

import "strconv"

func setBit(num int64, index uint, val bool) {
	decimal := strconv.FormatInt(num, 2)
	if val {
		[]rune(decimal)[index] = 1
	} else {
		[]rune(decimal)[index] = 0
	}
}
