package main

import "fmt"

const (
	wattToHPModifier = 1.3596
)

type Car struct {
	hp, weightInKilos float64
}

type GasEngine struct {
	price, hp float64
}

type ElecticEngine struct {
	watts, capacity, price float64
}

func main() {
	gas := GasEngine{
		price: 1000,
		hp:    300,
	}
	electric := ElecticEngine{
		watts:    400,
		capacity: 6000,
		price:    4000,
	}

	car1 := createCar(gas.hp, 2000)
	car2 := createCar(electric.castElectricToHP(), 4000)
	fmt.Println(car1)
	fmt.Println(car2)
}

func createCar(power, weight float64) Car {
	return Car{
		hp:            power,
		weightInKilos: weight,
	}
}

func (engine ElecticEngine) castElectricToHP() float64 {
	return engine.watts * wattToHPModifier
}
