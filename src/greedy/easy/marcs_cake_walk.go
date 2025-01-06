package easy

import (
	"math"
	"sort"
)

/*
 * calorie = [7 4 9 6]
 * Output: (2^0 * 9) + (2^1 * 7) + (2^2 * 6) + (2^3 * 4) = 79
 */
func MarcsCakewalk(calorie []int32) int64 {
	// Write your code here
	sort.Slice(calorie, func(i, j int) bool {
		return calorie[i] > calorie[j]
	})
	var rs int64
	for i, item := range calorie {
		rs += int64(item) * int64(math.Pow(2, float64(i)))
	}

	return rs
}
