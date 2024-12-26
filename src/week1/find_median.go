package main

func findMedian(nums []int) int {
	result := 0
	nums = bubleSort(nums)
	if len(nums)%2 == 0 {
		result = (nums[len(nums)/2] + nums[len(nums)/2-1]) / 2
	} else {
		result = nums[len(nums)/2]
	}
	return result
}

func bubleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
