package leetcode

import (
	"testing"
)

type LongestSubstringCase struct {
	str   string
	count int
}

var LongestSubstringCases = []LongestSubstringCase{
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
	for _, item := range LongestSubstringCases {
		result := lengthOfLongestSubstring(item.str)

		if result != item.count {
			t.Errorf("Longest substring of '%s' count: %d, expect: %d", item.str, result, item.count)
		}
	}
}
