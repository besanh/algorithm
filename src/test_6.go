package test6

/*
 * Complete the 'surfaceArea' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY A as parameter.
 */

func surfaceArea(A [][]int32) int32 {
	var sum int32 = 0
	H := int32(len(A))
	W := int32(len(A[0]))

	for i := int32(0); i < H; i++ {
		for j := int32(0); j < W; j++ {
			if A[i][j] > 0 {
				// Add top and bottom surfaces
				sum += 2

				// Add side surfaces
				// Check each direction: left, right, up, down
				// Left
				if j == 0 {
					sum += A[i][j] // No neighbor to the left
				} else {
					sum += max(0, A[i][j]-A[i][j-1])
				}

				// Right
				if j == W-1 {
					sum += A[i][j] // No neighbor to the right
				} else {
					sum += max(0, A[i][j]-A[i][j+1])
				}

				// Up
				if i == 0 {
					sum += A[i][j] // No neighbor above
				} else {
					sum += max(0, A[i][j]-A[i-1][j])
				}

				// Down
				if i == H-1 {
					sum += A[i][j] // No neighbor below
				} else {
					sum += max(0, A[i][j]-A[i+1][j])
				}
			}
		}
	}

	return sum
}

// Helper function to get maximum of two integers
func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
