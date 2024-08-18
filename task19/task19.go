package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	s := "Привет!"
	fmt.Println(revertString(s))
	s1 := ""
	fmt.Println(revertString(s1))
}

func revertString(s string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("Invalid string length")
	}
	casted := []rune(s)
	counter := len(casted) - 1
	var buf strings.Builder
	for counter >= 0 {
		buf.WriteRune(casted[counter])
		counter--
	}
	return buf.String(), nil
}
