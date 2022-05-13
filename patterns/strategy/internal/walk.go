package internal

import "fmt"

type WalkStrategy struct {
}

func (r *WalkStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 5
	total := endPoint - startPoint
	totalTime := total * 60
	fmt.Printf("Walk: Road A:[%d] to B:[%d] Avg speed:[%d] Total:[%d] Total time: [%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}
