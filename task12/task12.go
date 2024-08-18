package main

import "fmt"

func main() {
	stringsArr := []string{"cat", "cat", "dog", "tree", "cat"}
	setOfArr := make(map[string]struct{})

	for _, val := range stringsArr {
		if _, ok := setOfArr[val]; !ok {
			setOfArr[val] = struct{}{}
		}
	}

	for k, _ := range setOfArr {
		fmt.Println(k)
	}
}
