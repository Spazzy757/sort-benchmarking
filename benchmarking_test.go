package main

import (
	"sortbenchmarking/pkg/bubble"
	"sortbenchmarking/pkg/insertion"
	"sortbenchmarking/pkg/merge"
	"testing"
)

const arraySize = 10000

var arr = GenerateArray(arraySize)

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

func BenchmarkTestInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertion.Sort(arr)
	}
}
