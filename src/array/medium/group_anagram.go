package medium

import (
	"sort"
	"strings"
)

/*
 * Approach 1
 * Utilize hashmap and sort
 * Input: strs = ["eat","tea","tan","ate","nat","bat"]
 * Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 */
func GroupAnagrams(strs []string) [][]string {
	rs := make([][]string, 0)
	hashMap := make(map[string][]string)
	for _, item := range strs {
		runes := []rune(item)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		hashMap[string(runes)] = append(hashMap[string(runes)], item)
	}
	for _, item := range hashMap {
		rs = append(rs, item)
	}

	return rs
}

// Approach 2
func GroupAnagrams2(strs []string) [][]string {
	rs := make([][]string, 0)
	list := make(map[string][]string)
	for _, item := range strs {
		tmp := strings.Split(item, "")
		sort.Strings(tmp)
		sortedString := strings.Join(tmp, "")
		list[sortedString] = append(list[sortedString], item)
	}

	for _, val := range list {
		rs = append(rs, val)
	}

	return rs
}
