package internal

import "fmt"

type RoadStrategy struct {
}

func (r *RoadStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 30
	trafficJam := 2
	total := endPoint - startPoint
	totalTime := total * 40 * trafficJam
	fmt.Printf("Route: Road A:[%d] to B:[%d] Avg speed:[%d] Traffic jam:[%d] Total:[%d] Total time: [%d] min\n",
		startPoint, endPoint, avgSpeed, trafficJam, total, totalTime)
}
