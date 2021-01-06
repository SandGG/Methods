package main

import (
	"fmt"
)

type figure struct {
	x, y int
}

type values figure

func (f figure) area() int {
	return (f.x * f.y)
}

func area(f figure) int {
	return (f.x * f.y)
}

func (f values) valuesType() values {
	return f
}

func main() {
	var f = figure{x: 3, y: 4}
	fmt.Println(f.area())
	fmt.Println(area(f))

	var fT = values(f)
	fmt.Println(fT.valuesType())
}
