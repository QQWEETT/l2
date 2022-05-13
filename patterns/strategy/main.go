package main

import (
	"strategy/internal"
)

var (
	start      = 10
	end        = 100
	strategies = []internal.Strategy{
		&internal.PublicTransportStrategy{},
		&internal.RoadStrategy{},
		&internal.WalkStrategy{},
	}
)

func main() {
	nav := internal.Navigator{}
	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start, end)
	}
}
