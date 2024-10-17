package test1

import (
	"log"
	"math"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var resultA, resultB float64
	a := []int32{}
	b := []int32{}
	for i, item := range arr {
		a = append(a, item[i])
		j := len(arr[i]) - i - 1
		b = append(b, item[j])
		j--
	}
	for _, item := range a {
		resultA = float64(resultA + float64(item))
	}
	for _, item := range b {
		log.Println(item)
		resultB = float64(resultB + float64(item))
	}
	absRes := int32(math.Abs(resultA - resultB))
	return absRes
}
