package problems

import (
	"testing"
)

type P9Case struct {
	num  int
	want bool
}

var P9Cases = []P9Case{
	{1, true},
	{-123, false},
	{111, true},
	{121, true},
	{11, true},
	{123, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, item := range P9Cases {
		result := isPalindrome(item.num)

		if result != item.want {
			t.Errorf("Number %d is palindrome? expect: %v, result: %v", item.num, item.want, result)
		}
	}
}
