package main

import "fmt"

type Number interface {
	int | uint | uint8 | uint16 | uint32 | uint64 |
		float64 | float32
}

func main() {
	sl := []int{1, 5, 7, 9, 10}
	if index, ok := binarySearch(sl, 1); ok {
		fmt.Println(index)
	} else {
		fmt.Println("No such element in slice")
	}

}

func binarySearch[T Number](sl []T, val T) (int, bool) {
	left, right := 0, len(sl)-1
	for left <= right {
		index := (left + right) / 2
		if sl[index] == val {
			return index, true
		}
		if sl[index] > val {
			right = index - 1
		} else {
			left = index + 1
		}
	}
	return -1, false
}
