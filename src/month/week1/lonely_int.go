package week1

/**
 * a = [1, 2, 3, 4, 3, 2, 1]
 * Output: 4
 * I can use map or 2 pointers
 */
func Lonelyinteger(a []int32) int32 {
	// Write your code here
	hasInt := make(map[int32]int32)
	for _, num := range a {
		hasInt[num]++
	}

	var b int32
	for i, item := range hasInt {
		if item == 1 {
			b = i
			break
		}
	}
	return b
}
