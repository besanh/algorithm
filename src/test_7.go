package test7

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

// Function to count valid integers
func getTotalX(a []int32, b []int32) int32 {
	count := int32(0)
	for i := int32(1); i <= b[len(b)-1]; i++ {
		if calA(a, i) && calB(b, i) {
			count++
		}
	}
	return count
}

func calA(a []int32, i int32) bool {
	for _, item := range a {
		if i%item != 0 {
			return false
		}
	}
	return true
}

func calB(a []int32, i int32) bool {
	for _, item := range a {
		if item%i != 0 {
			return false
		}
	}
	return true
}
