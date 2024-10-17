package test13

import (
	"fmt"
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
	fmt.Println(2 ^ 3)
	total := int32(0)
	for i := a; i <= b; i++ {
		if math.Sqrt(float64(i)) == float64(int(math.Sqrt(float64(i)))) {
			total++
		}
	}
	return total
}
