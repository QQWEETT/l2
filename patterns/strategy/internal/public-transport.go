package internal

import "fmt"

type PublicTransportStrategy struct {
}

func (r *PublicTransportStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 60
	total := endPoint - startPoint
	totalTime := total * 40
	fmt.Printf("Public: Road A:[%d] to B:[%d] Avg speed:[%d] Total:[%d] Total time: [%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}
