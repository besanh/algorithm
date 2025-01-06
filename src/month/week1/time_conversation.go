package week1

import (
	"strconv"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */
// 12:01:01PM => 12:01:01
// 12:01:01AM => 00:01:01
func TimeConversion(s string) string {
	// Write your code here
	timeType := s[len(s)-2:]
	s = s[:len(s)-2]
	newString := strings.Split(s, ":")
	if timeType == "PM" {
		if newString[0] != "12" {
			tmp, _ := strconv.Atoi(newString[0])
			tmp += 12
			tmpInt := strconv.Itoa(tmp)
			newString[0] = tmpInt
		}

		s = strings.Join(newString, ":")
		return s
	} else if timeType == "AM" {
		if newString[0] == "12" {
			newString[0] = "00"
			s = strings.Join(newString, ":")
			return s
		} else {
			return s
		}
	}
	return s
}
