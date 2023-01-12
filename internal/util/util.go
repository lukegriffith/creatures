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
