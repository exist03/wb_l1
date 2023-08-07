package main

import (
	"fmt"
	"math"
	"wb_l1/24/point"
)

func main() {
	p1 := point.New()
	p2 := point.New()
	p1.SetX(0)
	p1.SetY(0)
	p2.SetX(5)
	p2.SetY(5)
	fmt.Println(interval(p1, p2))
}
func interval(p1, p2 *point.Point) float64 {
	return math.Sqrt(math.Pow(p1.GetX()-p2.GetX(), 2) + math.Pow(p1.GetY()-p2.GetY(), 2))
}
