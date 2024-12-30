package week1

import (
	"fmt"
	"sort"
)

/*
 * A = [2, 1, 3]
 * B = [7, 8, 9]
 * k = 10
 *
 * Permute of A and B in turn are A' and B'
 * A'[0] + B'[0] >= k
 * A'[1] + B'[1] >= k
 * A'[2] + B'[2] >= k
 *
 * Output: YES
 * If there is any value x < k => NO
 */
func TwoArrays(k int32, A []int32, B []int32) string {
	// Write your code here
	if len(A) != len(B) {
		return "NO"
	}

	sort.Slice(B, func(i, j int) bool {
		return B[i] > B[j]
	})

	list := []int{}
	for _, v := range A {
		list = append(list, int(v))
	}
	sort.Ints(list)
	fmt.Println("A:", list)
	fmt.Println("B: ", B)
	for i, item := range list {
		if item+int(B[i]) < int(k) {
			return "NO"
		}
	}
	return "YES"
}
