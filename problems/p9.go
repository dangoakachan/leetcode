package problems

// Problem 9. Palindrome Number
// URL: https://leetcode.com/problems/palindrome-number/description/

func isPalindrome(x int) bool {
	y, z := 0, x // the reverse integer

	if x < 0 {
		return false
	}

	for z > 0 {
		y = y*10 + z%10
		z /= 10
	}

	return x == y
}
