package pointers

import "sort"

/*
 * Input: people = [1,2], limit = 3
 * Output: 1 boat (1, 2)
 *
 * Input: people = [3,2,2,1], limit = 3
 * Output: 3 boats (1, 2), (2) and (3)
 *
 * Input: people = [3,5,3,4], limit = 5
 * Output: 4 boats (3), (3), (4), (5)
 * notme
 */
func RescueBoats(people []int, limit int) int {
	if len(people) == 0 || limit == 0 {
		return 0
	}

	sort.Ints(people)
	sum, left, right := 0, 0, len(people)-1
	for left <= right {
		if people[left]+people[right] <= limit {
			left++
		}
		right--
		sum++
	}

	return sum
}
