package internal

type Shape interface {
	GetType() string
	Accept(Visitor)
}
