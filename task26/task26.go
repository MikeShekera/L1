package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "Hello, world!"
	s1 := "asdA"
	s2 := "Привет!"
	fmt.Println(checkUniqueByMap(s))
	fmt.Println(checkUniqueByMap(s1))
	fmt.Println(checkUniqueByMap(s2))
}

// Выгоднее всего использовать map, так как нам извество количество элементов, которые содержит строка, а так же с использованием struct{} можно добиться оптимального использования памяти. Плюс поиск по map имеет константную сложность
func checkUniqueByMap(s string) bool {
	s = strings.ToLower(s)
	m := make(map[int32]struct{}, utf8.RuneCountInString(s))
	for _, v := range s {
		if _, ok := m[v]; ok {
			return false
		} else {
			m[v] = struct{}{}
		}
	}
	return true
}
