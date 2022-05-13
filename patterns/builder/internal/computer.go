package internal

import "fmt"

type Computer struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     bool
	GraphicCard int
}

func (pc *Computer) Print() {
	fmt.Printf("%s Core:[%d] Memory:[%d] GPU:[%d] Monitor: [%t]\n", pc.Brand, pc.Core, pc.Memory, pc.GraphicCard, pc.Monitor)
}
