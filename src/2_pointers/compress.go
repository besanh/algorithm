package pointers

import "strconv"

/*
 * Input: chars = ["a","a","b","b","c","c","c"]
 * Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
 * Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".
 * notme
 */
func Compress(chars []byte) int {
	n := len(chars)
	if n == 0 {
		return 0
	}

	write, i := 0, 0
	for i < n {
		j := i

		// Find the space
		for j < n && chars[i] == chars[j] {
			j++
		}

		// Update the count
		count := j - i
		chars[write] = chars[i]
		write++

		if count > 1 {
			countStr := strconv.Itoa(count)
			for _, v := range countStr {
				chars[write] = byte(v)
				write++
			}
		}

		// move pointer to next space
		i = j
	}

	return write
}
