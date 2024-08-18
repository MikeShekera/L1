package main

import "fmt"

func main() {
	var a, b int
	a = 13
	b = 7

	fmt.Println("a: ", a, ", b: ", b)

	a, b = b, a
	fmt.Println("a: ", a, ", b: ", b)
}
