package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	var b string
	var c bool
	ch := make(chan interface{})
	var d int64

	getTypeByReflect(a)
	getTypeByCast(b)
	getType(c)
	getTypeByCast(ch)
	getTypeByCast(d)
}

func getTypeByReflect(val interface{}) {
	fmt.Println(reflect.TypeOf(val))
}

func getTypeByCast(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan interface{}:
		fmt.Println("chan")
	default:
		fmt.Println("Can't get this type")
	}
}

func getType(val interface{}) {
	fmt.Println(fmt.Sprintf("%T", val))
}
