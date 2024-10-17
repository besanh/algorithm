package test8

/*
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	i := int32(0)
	for k := int32(0); k <= int32(len(s))-m; k++ {
		sum := int32(0)
		for j := int32(0); j < m; j++ {
			sum += s[k+j]
		}
		if sum == d {
			i++
		}
	}

	return i
}
