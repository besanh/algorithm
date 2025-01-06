package easy

import "sort"

/*
 * k = 3
 * contests = [[5, 1], [2, 1], [1, 1], [8, 1], [10, 0], [5, 0]]
 * 5 + 2 + 8 + 10 + 5 - 1 = 29
 */
func LuckBalance(k int32, contests [][]int32) int32 {
	// Write your code here
	var totalLucky int32
	importantLucky := make([]int32, 0)

	for _, item := range contests {
		if item[1] == 1 { // important
			importantLucky = append(importantLucky, item[0])
		} else {
			totalLucky += item[0]
		}
	}

	sort.Slice(importantLucky, func(i, j int) bool {
		return importantLucky[i] > importantLucky[j]
	})

	for i, luck := range importantLucky {
		if i < int(k) {
			totalLucky += luck
		} else {
			totalLucky -= luck
		}
	}

	return totalLucky

}
