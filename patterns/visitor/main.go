package main

import (
	"fmt"
	"visitor/internal"
)

func main() {
	square := &internal.Square{2}
	circle := &internal.Circle{3}
	rectangle := &internal.Rectangle{2, 3}

	areaCalculator := &internal.AreaCalculator{}
	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &internal.MiddleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)
}
