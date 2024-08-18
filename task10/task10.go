package main

import "fmt"

func main() {
	temperatureRange := []float32{-13, 24, -4.5, 15, 15.5, -25, -9, 5, 7, 2}
	sortedTemps := make(map[float32][]float32)
	for _, v := range temperatureRange {
		temp := int(v)
		temp /= 10
		temp *= 10
		if val, ok := sortedTemps[float32(temp)]; ok {
			val = append(val, v)
			sortedTemps[float32(temp)] = val
		} else {
			sortedTemps[float32(temp)] = []float32{v}
		}
	}

	fmt.Println(sortedTemps)
}
