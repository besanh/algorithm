package test12

/*
 * Complete the 'appendAndDelete' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING t
 *  3. INTEGER k
 */

func appendAndDelete(s string, t string, k int32) string {
	commonLength := getLongest(s, t)
	total := len(s) - len(commonLength) + len(t) - len(commonLength)
	if total <= int(k) {
		return "Yes"
	}
	return "No"
}

func getLongest(s, t string) string {
	minLength := len(t)
	if len(s) < len(t) {
		minLength = len(s)
	}

	for i := 0; i < minLength; i++ {
		if s[i] != t[i] {
			return s[:i]
		}
	}
	return s[:minLength]
}
