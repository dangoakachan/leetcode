package problems

import (
	"testing"
)

type P3Case struct {
	str   string
	count int
}

var P3Cases = []P3Case{
	{"abcabcbb", 3},
	{"abcacbe", 4},
	{"bbbb", 1},
	{"pwwkew", 3},
	{"b", 1},
	{"abcbb", 3},
	{"au", 2},
	{"dvdf", 3},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, item := range P3Cases {
		result := lengthOfLongestSubstring(item.str)

		if result != item.count {
			t.Errorf("Longest substring of '%s' count: %d, expect: %d", item.str, result, item.count)
		}
	}
}
