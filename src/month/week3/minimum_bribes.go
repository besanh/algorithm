package week3

import "fmt"

/*
 * The most important thing is check from right to left with 2 units
 * Neu 1 nguoi trong khoang i -> i-2 => return Chaotic
 * q = [2, 1, 5, 3, 4]
 * Output: 3
 * q = [2, 5, 1, 3, 4]
 * Output: Too chaotic
 */
func MinimumBribes(q []int32) {
	// Write your code here
	var bribes int32
	for i := len(q) - 1; i >= 0; i-- {
		if q[i]-int32(i-2) > 2 {
			fmt.Println("Too chaotic")
			return
		}

		for j := max(0, q[i]-2); j < int32(i); j++ {
			if q[j] > q[i] {
				bribes++
			}
		}
	}

	fmt.Println(bribes)
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
