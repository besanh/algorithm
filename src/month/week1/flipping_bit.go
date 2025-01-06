package week1

/**
 * Input: 3
 * 2147483647
 * 1
 * 0
 * Output:
 * 2147483648
 * 4294967294
 * 4294967295
 */
func FlippingBits(n int64) int64 {
	// Write your code here
	// Define the mask for a 32-bit unsigned integer
	// const max32Bit uint32 = 4294967295

	// Flip the bits using XOR with the mask
	flipped := ^uint32(n)

	return int64(flipped)
}
