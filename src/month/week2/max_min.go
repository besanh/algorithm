package week2

import "sort"

/*
 * Using slide window and greedy
 * k = 4
 * arr = [1, 2, 3, 4, 10, 20, 30, 40, 100, 200]
 * Output: max(1,2,3,4) - min(1,2,3,4) = 4 - 1 = 3
 */
func MaxMin(k int32, arr []int32) int32 {
	// Write your code here
	if k > int32(len(arr)) {
		return 0
	}

	list := []int{}
	for _, val := range arr {
		list = append(list, int(val))
	}

	sort.Ints(list)

	var rs int32
	var window []int32
	// Initial window
	for i := int32(0); i < k; i++ {
		window = append(window, int32(list[i]))
	}
	rs = window[len(window)-1] - window[0]

	// Slide the window with fixed size
	for i := k; i < int32(len(list)); i++ {
		window = append(window, int32(list[i]))
		window = window[1:]
		tmp := window[len(window)-1] - window[0]
		// Greedy
		if tmp < rs {
			rs = tmp
		}
	}

	return rs
}
