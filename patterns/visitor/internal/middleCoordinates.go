package internal

import "fmt"

type MiddleCoordinates struct {
	X int
	Y int
}

func (a *MiddleCoordinates) VisitForSquare(s *Square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) VisitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

func (a *MiddleCoordinates) VisitForrectangle(t *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
