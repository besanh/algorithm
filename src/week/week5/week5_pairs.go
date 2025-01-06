package week5

import (
	"sort"
)

/*
 * k = 2
 * arr = 1 5 3 4 2
 * output = 3([3,1] [4,2] [5,3])
 */
func Pair(k int32, arr []int32) int32 {
	// Write your code here
	list := []int{}
	for _, val := range arr {
		list = append(list, int(val))
	}
	sort.Ints(list)

	var rs int32
	slow, fast := 0, 1
	// Ensuring fast constantly greater than slow
	for fast < len(list) {
		diff := list[fast] - list[slow]

		if diff == int(k) {
			rs++
			slow++
			fast++
		} else if diff < int(k) {
			fast++
		} else {
			slow++
		}

		// Ensuring slow always less than fast
		if slow == fast {
			fast++
		}
	}

	return rs
}
