package test14

import (
	"math"
)

/*
 * Complete the 'squares' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER a
 *  2. INTEGER b
 */
func squares(a int32, b int32) int32 {
	// Write your code here
	// Calculate the smallest integer greater than or equal to the square root of a
	start := int32(math.Ceil(math.Sqrt(float64(a))))

	// Calculate the largest integer less than or equal to the square root of b
	end := int32(math.Floor(math.Sqrt(float64(b))))

	// The count of perfect squares is the number of integers from start to end
	if start > end {
		return 0
	}
	return end - start + 1
}
