package week2

/*
 * Complete the 'flippingMatrix' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY matrix as parameter.
 */
/**
* Get 4 corners by moving the nxn matrix
 */
func FlippingMatrix(matrix [][]int32) int32 {
	// Write your code here
	// Chieu dai cua matrix con nam ben trai
	n := len(matrix) / 2
	maxSum := 0

	// 2 loops duyet qua tung phan tu cua matrix con
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rs := maxPosition(
				// Top left
				int(matrix[i][j]),
				// Top right
				int(matrix[i][2*n-j-1]),
				// Bot left
				int(matrix[2*n-i-1][j]),
				// Bot right
				int(matrix[2*n-i-1][2*n-j-1]),
			)
			maxSum += rs
		}
	}
	return int32(maxSum)
}

// 4 gia tri doi xung
func maxPosition(a, b, c, d int) int {
	maximum := a
	if maximum < b {
		maximum = b
	}
	if maximum < c {
		maximum = c
	}
	if maximum < d {
		maximum = d
	}

	return maximum
}
