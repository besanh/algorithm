package week1

func CountingSort(arr []int32) []int32 {
	// Write your code heres
	count := make([]int32, 100)

	for _, val := range arr {
		count[val]++
	}

	return count
}
