package week1

/*
 * Input: We promptly judged antique ivory buckles for the next prize => pangram
 * Input: We promptly judged antique ivory buckles for the prize => not pangram
 */
func Pangrams(s string) string {
	// Write your code here
	list := make(map[rune]bool)
	for _, char := range s {
		// Convert to lowercase if it is a letter
		if char >= 'A' && char <= 'Z' {
			char += 'a' - 'A'
		}
		if char >= 'a' && char <= 'z' {
			list[char] = true
		}
	}
	if len(list) == 26 {
		return "pangram"
	}
	return "not pangram"
}
