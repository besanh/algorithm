package week3

func PalindromeIndex(s string) int32 {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			if isParalindrome(s, left+1, right) {
				return int32(left)
			} else if isParalindrome(s, left, right-1) {
				return int32(right)
			}
			return -1
		}
	}
	return -1
}

func isParalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
