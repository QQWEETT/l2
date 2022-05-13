package internal

type Strategy interface {
	Route(startPoint int, endPoint int)
}
