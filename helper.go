package main

import (
	"math/rand"
)

func GenerateArray(n int) []int {
	// we set the seed to 1 to have the same array generated
	r := rand.New(rand.NewSource(1))
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, r.Intn(10000000000))
	}
	return result
}
