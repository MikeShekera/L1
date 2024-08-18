package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sl = remove(sl, 1)
	fmt.Println(sl)
	sl = fastRemove(sl, 1)
	fmt.Println(sl)
}

// Медленное удаление с сохранением порядка элементов
func remove(sl []int, removeIndex uint) []int {
	if len(sl) > 1 {
		return append(sl[:removeIndex], sl[removeIndex+1:]...)
	} else {
		return make([]int, 0)
	}
}

// Быстрое удаление. Применимо только в том случае, если нет цели сохранить исходный порядок элементов
func fastRemove(sl []int, removeIndex int) []int {
	if len(sl) > 1 {
		sl[removeIndex] = sl[len(sl)-1]
		return sl[:len(sl)-1]
	} else {
		return make([]int, 0)
	}
}
