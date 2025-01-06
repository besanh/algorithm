package week4

func SuperDigit(n string, k int32) int32 {
	// Write your code here
	// Approach 1
	// var backtrack func(num string) int32
	// backtrack = func(num string) int32 {
	// 	if len(num) == 1 {
	// 		// Convert string to int
	// 		return int32(num[0] - '0')
	// 	}

	// 	var sum int32
	// 	for _, char := range num {
	// 		sum += int32(char - '0')
	// 	}

	// 	return backtrack(fmt.Sprint(sum))
	// }

	// // The number n is created by concatenating the string k times so the initial
	// initialString := strings.Repeat(n, int(k))
	// return backtrack(initialString)

	// Approach 2: using recursive
	var digiSum int64
	for _, char := range n {
		digiSum += int64(char - '0')
	}

	newString := digiSum * int64(k)

	var computeSuperDigit func(num int64) int32
	computeSuperDigit = func(i int64) int32 {
		if i < 10 {
			return int32(i)
		}

		var sum int64
		for i > 0 {
			sum += i % 10 // Add last digit
			i /= 10       // Remove last digit
		}

		return computeSuperDigit(sum)
	}

	return computeSuperDigit(newString)
}
