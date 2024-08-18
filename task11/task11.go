package main

import "fmt"

func main() {
	m1 := make(map[interface{}]struct{})
	m2 := make(map[interface{}]struct{})

	m1["data"] = struct{}{}
	m1[125] = struct{}{}
	m1["mystring"] = struct{}{}

	m2["data1"] = struct{}{}
	m2[125] = struct{}{}
	m2["mystring"] = struct{}{}

	var intersectValues map[interface{}]struct{}
	if len(m1) < len(m2) {
		intersectValues = getIntersectValues(m1, m2)
	} else {
		intersectValues = getIntersectValues(m2, m1)
	}
	for k, _ := range intersectValues {
		fmt.Println(k)
	}
}

func getIntersectValues(shortMap map[interface{}]struct{}, longMap map[interface{}]struct{}) map[interface{}]struct{} {
	intersectValues := make(map[interface{}]struct{})
	for k, _ := range shortMap {
		if _, ok := longMap[k]; ok {
			intersectValues[k] = struct{}{}
		}
	}
	return intersectValues
}
