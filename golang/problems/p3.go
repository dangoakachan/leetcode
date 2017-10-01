package problems

// Problem 3. Longest Substring Without Repeating Characters
// URL: https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	n := len(s)

	if n <= 1 {
		return n
	}

	tbl := make(map[byte]int, n)
	lo, longest, curr := 0, 0, 0

	for hi := 0; hi < n; hi++ {
		if k, ok := tbl[s[hi]]; ok && k >= lo {
			lo = k + 1
		}

		tbl[s[hi]] = hi
		curr = hi - lo + 1

		if curr > longest {
			longest = curr
		}
	}

	return longest
}
