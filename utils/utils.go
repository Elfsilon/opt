package utils

import (
	"log"
	"math/rand"
	"time"
)

// RandFloat ...
func RandFloat(min, max float64) float64 {
	rnd := rand.Float64()*(max-min) + min
	return rnd
}

// RandInt ...
func RandInt(min, max int) int {
	rnd := rand.Int()*(max-min) + min
	return rnd
}

// Contains ...
func Contains(item string, slice []string) bool {
	for _, str := range slice {
		if str == item {
			return true
		}
	}
	return false
}

// TimeTrack ...
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s\n", name, elapsed)
}
