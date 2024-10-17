package test10

/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	var result int32
	var max int32
	tmp := map[int32]int32{}
	for _, item := range arr {
		tmp[item]++
	}

	for k, item := range tmp {
		if item > max || (item == max && k < result) {
			max = item
			result = k
		}
	}
	return result
}

func exist(item int32, slices []int32) bool {
	for i := int32(0); i < int32(len(slices)); i++ {
		if item == slices[i] {
			return true
		}
	}
	return false
}
