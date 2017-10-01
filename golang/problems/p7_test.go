package problems

import (
	"testing"
)

type P7Case struct {
	num         int
	reverse_num int
}

var P7Cases = []P7Case{
	{123, 321},
	{-123, -321},
	{10, 1},
	{100, 1},
	{1000000003, 0},
}

func TestReverse(t *testing.T) {
	for _, item := range P7Cases {
		result := reverse(item.num)

		if result != item.reverse_num {
			t.Errorf("Reverse number: %d, expect: %d", result, item.reverse_num)
		}
	}
}
