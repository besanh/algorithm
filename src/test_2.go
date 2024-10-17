package test2

import (
	"fmt"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	var totalPositive, totalNegative, totalZero float64
	for _, item := range arr {
		if item == 0 {
			totalZero++
		} else if item > 0 {
			totalPositive++
		} else {
			totalNegative++
		}
	}

	fmt.Printf("%.6f\n%.6f\n%.6f", float64(totalPositive/float64(len(arr))), float64(totalNegative/float64(len(arr))), float64(totalZero/float64(len(arr))))
	return
}
