package test4

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	var min, max, el int64
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			min = int64(arr[i])
		} else if i == len(arr)-1 {
			max = int64(arr[i])
		} else {
			el += int64(arr[i])
		}
	}
	fmt.Print(min+el, max+el)
}
