package util

import (
	"math/rand"
	"time"
)

func RandomInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func RandomFloat(min int, max int) float64 {
	rand.Seed(time.Now().UnixNano())
	return float64(rand.Intn(max-min+1) + min)
}

func MinMax(array []float64) (int, int) {
	var max float64 = array[0]
	var maxIndex int
	var min float64 = array[0]
	var minIndex int
	for i, value := range array {
		if max < value {
			max = value
			maxIndex = i
		}
		if min > value {
			min = value
			minIndex = i
		}
	}
	return minIndex, maxIndex
}
