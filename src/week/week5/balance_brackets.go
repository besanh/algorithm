package week5

/**
 * Sample input
 * {[()]}
 * {[(])}
 * {{[[(())]]}}
 */
func IsBalanced(s string) string {
	// Write your code here
	stack := []rune{}
	matchBrackets := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, char := range s {
		if char == '{' || char == '(' || char == '[' {
			stack = append(stack, char)
		} else if char == '}' || char == ')' || char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != matchBrackets[char] {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return "NO"
	}

	return "YES"
}
