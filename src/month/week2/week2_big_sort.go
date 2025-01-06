package week2

import "sort"

/*
 * unsorted: 6, 31415926535897932384626433832795, 1, 3, 10, 3, 5
 * sorted: 1, 3, 3, 5, 6, 10, 31415926535897932384626433832795
 */
func BigSorting(unsorted []string) []string {
	// Write your code here
	sort.Slice(unsorted, func(i, j int) bool {
		if len(unsorted[i]) != len(unsorted[j]) {
			return len(unsorted[i]) < len(unsorted[j])
		}

		return unsorted[i] < unsorted[j]
	})

	return unsorted
}
