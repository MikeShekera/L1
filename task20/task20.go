package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Мама мыла раму"
	s1 := "Привет!"
	fmt.Println(revertPhrase(s))
	fmt.Println(revertPhrase(s1))
}

func revertPhrase(s string) string {
	str := strings.Split(s, " ")
	if len(str) > 1 {
		var buf strings.Builder
		counter := len(str) - 1
		for counter >= 0 {
			buf.WriteString(str[counter])
			buf.WriteString(" ")
			counter--
		}
		return buf.String()
	}
	return s
}
