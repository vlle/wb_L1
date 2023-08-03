package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x, y int
}

func NewPoint() Point {
	return Point{}
}

func NewPointXY(x, y int) Point {
	return Point{x: x, y: y}
}

func (p1 Point) Distance(p2 Point) float64 {
	if p1.x == p2.x {
		return math.Abs(float64(p1.y - p2.y))
	} else if p1.y == p2.y {
		return math.Abs(float64(p1.x - p2.x))
	} else {
		a := math.Abs(float64(p1.x - p2.x))
		b := math.Abs(float64(p1.y - p2.y))
		return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	}
}

func t24() {
	p1 := NewPointXY(1, 4)
	p2 := NewPointXY(4, 4)
	fmt.Println(p1.Distance(p2))
}
