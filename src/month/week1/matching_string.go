package week1

import "regexp"

/**
 * strings = ["ab", "ab", "abc"]
 * queries = ["ab", "abc", "bc"]
 * @return [2, 1, 0]
 */
func MatchingStrings(strings, queries []string) []int32 {
	// Write your code here
	result := []int32{}
	for _, val := range queries {
		count := 0
		for _, v := range strings {
			if match, _ := regexp.MatchString(v, val); match {
				count++
			}
		}
		result = append(result, int32(count))
		count = 0
	}
	return result
}
