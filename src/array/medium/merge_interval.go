package medium

import "sort"

/*
 * Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 */
func MergeIntervals(intervals [][]int) [][]int {
	// Write your code here
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	rs := [][]int{}
	current := intervals[0]
	for _, item := range intervals[1:] {
		if item[0] < current[1] {
			if item[1] < current[1] {
				current[1] = item[1]
			}
		} else {
			rs = append(rs, current)
			current = item
		}
	}

	rs = append(rs, current)

	return rs
}
