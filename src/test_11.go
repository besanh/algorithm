package test11

import (
	"fmt"
	"math/big"
)

/*
 * Complete the 'extraLongFactorials' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func extraLongFactorials(n int32) {
	// Write your code here
	result := big.NewInt(1)
	for k := int64(1); k <= int64(n); k++ {
		result.Mul(result, big.NewInt(k))
	}
	fmt.Println(result.String())
}
