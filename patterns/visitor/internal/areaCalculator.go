package internal

import (
	"fmt"
)

type AreaCalculator struct {
	Area int
}

func (a *AreaCalculator) VisitForSquare(s *Square) {
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) VisitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *AreaCalculator) VisitForrectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}
