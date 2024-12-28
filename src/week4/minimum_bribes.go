package week4

import "fmt"

/**
* Moi nguoi co the bribe(hoi lo) 2 lan, neu nhieu hon 2 lan => print Chaostic
* Da di chuyen roi thi khong the lui lai
* Tinh tong so lan bribe can thiet de dua q ve ban dau(sorted array) => print Sum
 */
func MinimumBribes(q []int32) {
	// Write your code here
	sum := 0
	for i := len(q) - 1; i >= 0; i-- {
		// Check the previous element, if it is more than i+2 => Chaostic
		// Why do I use i+1 ? => Because index start from 0
		if q[i] > int32(i+1)+2 {
			fmt.Println("Too chaotic")
			return
		}

		// Permutation
		for j := max(0, int(q[i]-2)); j < i; j++ {
			if q[j] > q[i] {
				sum++
			}
		}
	}

	fmt.Println(sum)
}
