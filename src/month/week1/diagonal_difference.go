package week1

import "math"

/*
 * 1 2 3
 * 4 5 6
 * 9 8 9
 * left = 1 + 5 + 9 = 15
 * right = 3 + 5 + 9 = 17
 * output = |15-17| = 2
 */
func DiagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var left, right int32
	for i := int32(0); i < int32(len(arr)); i++ {
		left += arr[i][i]
		right += arr[i][int32(len(arr[i]))-1-i]
	}

	return int32(math.Abs(float64(left - right)))
}
