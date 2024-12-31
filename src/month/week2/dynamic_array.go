package week2

/*
 * n = 2 length = 5
 * 1 0 5
 * 1 1 7
 * 1 0 3
 * 2 1 0
 * 2 1 1
 * Output: [7, 3]
 * notme
 */
func DynamicArray(n int32, queries [][]int32) []int32 {
	// Write your code here
	seqList := make([][]int32, n)
	var rs []int32
	var lastAnswer int32

	for _, val := range queries {
		queryType, x, y := val[0], val[1], val[2]
		idx := (x ^ lastAnswer) % n

		// Insert
		if queryType == 1 {
			seqList[idx] = append(seqList[idx], y)
		} else if queryType == 2 {
			// Query and insert
			tmp := seqList[idx]
			lastAnswer = tmp[y%int32(len(tmp))]
			rs = append(rs, lastAnswer)
		}
	}

	return rs
}
