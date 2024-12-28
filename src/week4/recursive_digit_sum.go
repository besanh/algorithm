package week4

import (
	"fmt"
	"strings"
)

func SuperDigit(n string, k int32) int32 {
	// Write your code here
	var backtrack func(num string) int32
	backtrack = func(num string) int32 {
		if len(num) == 1 {
			// Convert string to int
			return int32(num[0] - '0')
		}

		var sum int32
		for _, char := range num {
			sum += int32(char - '0')
		}

		return backtrack(fmt.Sprint(sum))
	}

	// The number n is created by concatenating the string k times so the initial
	initialString := strings.Repeat(n, int(k))
	return backtrack(initialString)
}
