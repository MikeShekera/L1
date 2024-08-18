package main

import (
	"fmt"
	"math/big"
)

// Вариант с использованием math.big. В таком контексте число может иметь произвольную длину
func main() {
	a := new(big.Int)
	b := new(big.Int)
	temp := new(big.Int)
	a.SetString("452538785667345", 10)
	b.SetString("346548443563467567897", 10)

	fmt.Println(sum(a, b, temp))
	fmt.Println(substract(a, b, temp))
	fmt.Println(multiply(a, b, temp))
	fmt.Println(divide(a, b, temp))

}

func sum(a, b, temp *big.Int) *big.Int {
	return temp.Add(a, b)
}

func substract(a, b, temp *big.Int) *big.Int {
	return temp.Sub(a, b)
}

func multiply(a, b, temp *big.Int) *big.Int {
	return temp.Mul(a, b)
}

func divide(a, b, temp *big.Int) *big.Int {
	if len(b.Bits()) != 0 {
		return temp.Div(a, b)
	} else {
		fmt.Println("Can't divide by zero")
		return nil
	}
}

//Вариант с использованием базовых типов. Может работать с числами больше, чем 2^20, но есть верхняя граница значений, которой является максимально допустимая длина int64
/*func main() {
	var a, b int32
	a = math.MaxInt32
	b = math.MaxInt32
	fmt.Println(sum(a, b))
	fmt.Println(substract(a, b))
	fmt.Println(multuply(a, b))
	fmt.Println(divide(a, b))
}

func sum(a, b int32) int64 {
	return int64(a) + int64(b)
}

func substract(a, b int32) int64 {
	return int64(a) - int64(b)
}

func multuply(a, b int32) int64 {
	return int64(a) * int64(b)
}

func divide(a, b int32) int32 {
	return a / b
}*/
