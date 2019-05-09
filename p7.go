package leetcode

// Problem 7. Reverse Integer
// URL: https://leetcode.com/problems/reverse-integer/description/

const (
	MaxInt32 int = 1<<31 - 1
	MinInt32     = -MaxInt32 - 1
)

func reverse(x int) int {
	y := 0    // the reverse integer
	sign := 1 // 1: positive, -1: negative

	if x < 0 {
		sign = -1
		x *= -1
	}

	// x is a 32-bit number, so do not need check overflow here
	for x > 0 {
		y = y*10 + x%10
		x /= 10
	}

	y = sign * y

	// We check the overflow here now
	if y > MaxInt32 || y < MinInt32 {
		y = 0
	}

	return y
}
