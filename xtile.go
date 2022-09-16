package xtile

import "math"

// ZoomForBounds could determine which zoom level should be used.
//
// copy from https://stackoverflow.com/questions/4266754
func ZoomForBounds(lngWest float64, latSouth float64, lngEat float64, latNorth float64) int {
	latDiff := math.Abs(latNorth - latSouth)
	lngDiff := math.Abs(lngWest - lngEat)
	maxDiff := math.Max(latDiff, lngDiff)
	if maxDiff < 360/math.Pow(2, 20) {
		return 21
	}
	zoom := int((-1 * ((math.Log(maxDiff) / math.Log(2)) - (math.Log(360) / math.Log(2)))))
	if zoom < 1 {
		zoom = 1
	}
	return zoom
}
