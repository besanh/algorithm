package medium

import "strconv"

/*
 * Problem: https://leetcode.com/problems/evaluate-reverse-polish-notation/description/?envType=problem-list-v2&envId=array
 * Time complexity: O(n) because I visit each element once
 * Space complexcity: O(n) because I use stack
 * DSA: Stack
 * My solution: use stack, easy to implement
 * Explaination: I iterate each item in array and check if it is a number, if yes, I push it to stack,
 * if it is an operator, I pop two elements from stack, calculate the result and push the result back to stack,
 * repeat until all elements in array are processed
 */
func EvalRPN(tokens []string) int {
	// Write your code here
	stack := []int{}
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]

			// remove last two elements from stack
			stack = stack[:len(stack)-2]
			result := 0
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}
			stack = append(stack, result)

		default:
			val, err := strconv.Atoi(token)
			if err != nil {
				panic("invalid token: " + token)
			}
			stack = append(stack, val)
		}
	}

	return stack[0]
}
