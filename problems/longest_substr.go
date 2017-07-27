package leetcode

func lengthOfLongestSubstring(s string) int {
	n := len(s)

	if n <= 1 {
		return n
	}

	tbl := make(map[rune]int, n)
	lo, max := 0, 0

	for i, c := range s {
		if j, ok := tbl[c]; ok && j >= lo {
			if i-lo > max {
				max = i - lo
			}

			lo = j + 1
		}

		tbl[c] = i
	}

	if n-lo > max {
		max = n - lo
	}

	return max
}
