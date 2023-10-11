package main

import (
	"math/rand"
	"sortbenchmarking/pkg/bubble"
	"sortbenchmarking/pkg/merge"
	"testing"
	"time"
)

const arraySize = 10000

func generateArray(n int) []int {
	rand.Seed(time.Now().UnixNano())
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, rand.Intn(10000000000))
	}
	return result
}

var arr = generateArray(arraySize)

func BenchmarkTestBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubble.Sort(arr)
	}
}

func BenchmarkTestMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		merge.Sort(arr)
	}
}
