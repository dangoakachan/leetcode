package leetcode

import (
	"testing"
)

// P1Cases is a list of test cases for problem
type P1Case struct {
	arr    []int
	target int
	want   []int
}

// P1Cases is a list of test cases for problem
var P1Cases = []P1Case{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{2, 7, 11, 15}, 11, nil},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{2}, 2, nil},
}

func TestTwoSum(t *testing.T) {
	for _, item := range P1Cases {
		result := twoSum(item.arr, item.target)

		if !CompareSlice(result, item.want) {
			t.Errorf("Two sum result: %v, expect: %v", result, item.want)
		}
	}
}

func TestTwoSum1(t *testing.T) {
	for _, item := range P1Cases {
		result := twoSum1(item.arr, item.target)

		if !CompareSlice(result, item.want) {
			t.Errorf("Two sum result: %v, expect: %v", result, item.want)
		}
	}
}
