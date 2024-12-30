package week2

func PageCount(n int32, p int32) int32 {
	// Write your code here
	var frontLip, backLip int32
	frontLip = p / 2
	backLip = (n / 2) - (p / 2)
	if frontLip < backLip {
		return frontLip
	}
	return backLip
}
