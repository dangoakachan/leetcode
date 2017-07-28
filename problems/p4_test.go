package problems

import (
	"testing"
)

type P4Case struct {
	nums1  []int
	nums2  []int
	median float64
}

var P4Cases = []P4Case{
	{[]int{1, 3}, []int{2}, 2.0},
	{[]int{1, 2}, []int{3, 4}, 2.5},
	{[]int{}, []int{1, 2, 3}, 2.0},
	{[]int{1, 2, 3, 4}, []int{}, 2.5},
	{[]int{}, []int{}, 0},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for _, item := range P4Cases {
		result := findMedianSortedArrays(item.nums1, item.nums2)

		if result != item.median {
			t.Errorf("Find median: %f, expect: %f\n", result, item.median)
		}
	}
}
