package main

import (
	"fmt"
	"math"
)

func getKey(value float64) int {
	return int(value - math.Mod(value, 10))
}

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	storage := make(map[int][]float64)
	for _, v := range arr {
		key := getKey(v)
		if _, ok := storage[key]; !ok {
			storage[key] = make([]float64, 0)
		}
		storage[key] = append(storage[key], v)
	}
	fmt.Println(storage)
}
