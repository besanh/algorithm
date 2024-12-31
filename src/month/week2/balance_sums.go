package week2

/*
 * The most importance is that the sum of the left part of the array is equal to the sum of the right part of the array
 * leftSum = rightSum
 * arr = 1 2 3 => NO
 * arr = 1 2 3 3 => YES
 */
func BalanceSums(arr []int32) string {
	// Write your code here
	var totalSum int32
	for _, val := range arr {
		totalSum += val
	}

	var leftSum int32
	for _, val := range arr {
		if leftSum == totalSum-leftSum-val {
			return "YES"
		}
		leftSum += val
	}

	return "NO"
}
