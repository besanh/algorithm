package test3

import (
	"fmt"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	// Write your code here
	hashtag := "#"
	space := " "
	for i := 0; i < int(n); i++ {
		for j := int(n); j >= 0; j-- {
			if j == 0 {
				fmt.Print("\n")
			} else if j-1 <= i {
				fmt.Print(hashtag)
			} else {
				fmt.Printf(space)
			}
		}
	}
}
