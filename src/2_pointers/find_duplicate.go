package pointers

import "sort"

/*
 * [1,3,4,2,2] => 2
 */
func FindDuplicate(nums []int) int {
	sort.Ints(nums)
	slow, fast := 0, 1

	for slow != fast {
		if nums[slow] == nums[fast] {
			return nums[slow]
		}
		fast++
		slow++
	}

	return -1
}
