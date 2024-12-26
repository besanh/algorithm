package week3

import (
	"sort"
)

func FindZigZagSequence(arr []int) []int {
	// 2, 3, 5, 1, 4, 6, 7 => 1 2 3 7 6 5 4
	sort.Ints(arr)
	mid := len(arr) / 2
	arr[mid], arr[len(arr)-1] = arr[len(arr)-1], arr[mid]
	sort.Sort(sort.Reverse(sort.IntSlice(arr[mid:])))
	return arr
}
