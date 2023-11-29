package main

import (
	"flag"
	"fmt"
	"sortbenchmarking/pkg/bubble"
	"sortbenchmarking/pkg/insertion"
	"sortbenchmarking/pkg/merge"
)

var (
	sortType   string
	listLength int
)

func main() {
	flag.StringVar(&sortType, "type", "merge", "sorting algorithm")
	flag.IntVar(&listLength, "list-length", 1000000, "length of list")

	flag.Parse()

	var arr = GenerateArray(listLength)
	var ans []int
	switch sortType {
	case "bubble":
		ans = bubble.Sort(arr)
	case "merge":
		ans = merge.Sort(arr)
	case "insertion":
		ans = insertion.Sort(arr)
	default:
		panic("sort type unkown")
	}
	fmt.Printf(
		"list of length %v was sorted using %s sort",
		len(ans),
		sortType,
	)
}
