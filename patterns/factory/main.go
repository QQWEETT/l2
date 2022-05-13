package main

import (
	"factory2/internal"
)

var types = []string{
	internal.PersonalComputerType,
	internal.NotebookType,
	internal.ServerType,
}

func main() {
	for _, typeName := range types {
		computer := internal.New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}
}
