package problems

// Problem 3. Longest Substring Without Repeating Characters
// URL: https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	n := len(s)

	if n <= 1 {
		return n
	}

	tbl := make(map[byte]int, n)
	lo, largest := -1, 0

	for i := 0; i < n; i++ {
		if j, ok := tbl[s[i]]; ok && j > lo {
			lo = j
		}

		tbl[s[i]] = i

		if i-lo > largest {
			largest = i - lo
		}
	}

	return largest
}
