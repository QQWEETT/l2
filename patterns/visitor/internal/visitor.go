package internal

// Добавляем функциональность
type Visitor interface {
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
	VisitForrectangle(*Rectangle)
}
