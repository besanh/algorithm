package week2

/*
 * n = 7
 * ar = [1, 2, 1, 2, 1, 3, 2]
 * [1, 1], [2, 2]
 * Output: 2
 */
func SockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	hashMap := make(map[int32]int32, 0)
	for _, item := range ar {
		hashMap[item]++
	}

	var rs int32
	for _, val := range hashMap {
		rs += val / 2
	}

	return rs
}
