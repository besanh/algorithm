package medium

import (
	"sort"
)

/*
 * hIndex = citations[h]≥h+1
 * citations = [3,0,6,1,5]
 * Output: 3
 */
func HIndex(citations []int) int {
	rs := 0
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	for i, val := range citations {
		if val >= i+1 {
			rs = i + 1
		} else {
			break
		}
	}

	return rs
}
