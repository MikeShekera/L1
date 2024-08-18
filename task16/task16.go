package main

import "fmt"

type Number interface {
	int | uint | uint8 | uint16 | uint32 | uint64 |
	float64 | float32
}

func main() {
	/*sl := []int{8, 3, 5, 1, 9, 6}
	fmt.Println(sort(sl))*/
	sl1 := []int{8, 3, 5, 1, 9, 6, 0, 11, 4}
	noAppendSort(sl1)
	fmt.Println(sl1)
}

func sort[T Number](sl []T) []T {
	if len(sl) <= 1 {
		return sl
	}
	var left, right []T
	control := sl[len(sl)-1]
	for _, v := range sl {
		if v < control {
			left = append(left, v)
		}
		if v > control {
			right = append(right, v)
		}
	}
	end := append(sort(left), control)
	return append(end, sort(right)...)
}

func noAppendSort[T Number](sl []T) {
	if len(sl) <= 1 {
		return
	}
	control := sl[len(sl)-1]
	counter := 0
	for i, v := range sl {
		if v < control {
			sl[i], sl[counter] = sl[counter], sl[i]
			counter++
		}
	}
	sl[(len(sl)-1)], sl[counter] = sl[counter], sl[(len(sl)-1)]
	noAppendSort(sl[:counter])
	noAppendSort(sl[counter+1:])
}
