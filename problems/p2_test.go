package problems

import (
	"testing"

	"github.com/dangoakachan/leetcode/common"
)

// P2Case defines the single test case for problem
type P2Case struct {
	l1   []int
	l2   []int
	want []int
}

// P2Cases is a list of test cases for problem
var P2Cases = []P2Case{
	{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
	{[]int{2, 4, 3}, []int{5, 6, 4, 5}, []int{7, 0, 8, 5}},
	{[]int{}, []int{}, nil},
	{[]int{5}, []int{5}, []int{0, 1}},
	{nil, nil, nil},
}

// Generates a list from the input array
func makeList(arr []int) *ListNode {
	head := &ListNode{0, nil}
	tail := head

	for _, v := range arr {
		tail.Next = &ListNode{v, nil}
		tail = tail.Next
	}

	return head.Next
}

// Converts the list to an array
func convertToSlice(l1 *ListNode) []int {
	var arr []int

	for l1 != nil {
		arr = append(arr, l1.Val)
		l1 = l1.Next
	}

	return arr
}

func TestAddTwoNumbers(t *testing.T) {
	for _, item := range P2Cases {
		result := convertToSlice(addTwoNumbers(makeList(item.l1), makeList(item.l2)))

		if !common.CompareSlice(result, item.want) {
			t.Errorf("Add two numbers result: %v, expect: %v", result, item.want)
		}
	}
}
