package test5

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var max int32 = candles[0]
	var count int32 = 0

	for _, candle := range candles {
		if candle > max {
			max = candle
			count = 1
		} else if candle == max {
			count++
		}
	}

	return count
}
