package leetcode

import (
	"testing"
)

type TwoSumCase struct {
	arr    []int
	target int
	want   []int
}

var TwoSumCases = []TwoSumCase{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{2, 7, 11, 15}, 11, nil},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{2}, 2, nil},
}

func TestTwoSum(t *testing.T) {
	for _, item := range TwoSumCases {
		result := twoSum(item.arr, item.target)

		if !compareSlice(result, item.want) {
			t.Errorf("Two sum result: %v, expect: %v", result, item.want)
		}
	}
}

func TestTwoSum1(t *testing.T) {
	for _, item := range TwoSumCases {
		result := twoSum1(item.arr, item.target)

		if !compareSlice(result, item.want) {
			t.Errorf("Two sum result: %v, expect: %v", result, item.want)
		}
	}
}
