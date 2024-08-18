package main

import (
	"fmt"
	"math"
)

type Human struct {
	height, weight float64
	age            int
	name           string
}

type Action struct {
	Human
}

func (h *Human) GetBMI() float64 {
	return h.weight / math.Pow(h.height, 2)
}

func (a *Action) GetBMI() string {
	return fmt.Sprintf("BMI of %s is %v", a.name, a.Human.GetBMI())
}

func main() {
	a := Action{
		Human{
			height: 1.7,
			weight: 62,
			name:   "Billy",
		}}

	fmt.Println(a.GetBMI())
}
