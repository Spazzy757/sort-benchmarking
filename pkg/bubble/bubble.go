package bubble

func Sort(arr []int) []int {
	// First Loop Over the Array
	for i := 0; i < len(arr)-1; i++ {
		// Second Loop Over the Array
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
