package medium

/*
 * Find elements that appear more than n/3 times
 * Input: nums = [3,2,3]
 * Output: [3]
 */
func MajorityElement(nums []int) []int {
	rs := make([]int, 0)
	point := len(nums) / 3
	hashMap := make(map[int]int, 0)
	for _, item := range nums {
		hashMap[item]++
	}
	for num, val := range hashMap {
		if val > point {
			rs = append(rs, num)
		}
	}

	return rs
}
