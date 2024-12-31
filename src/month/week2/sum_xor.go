package week2

func SumXor(n int64) int64 {
	// Write your code here
	var rs int64
	for i := int64(0); i <= n; i++ {
		if n+i == n^i {
			rs++
		}
	}

	return rs
}
