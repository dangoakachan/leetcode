package leetcode

import (
	"testing"
)

type AddTwoNumbersCase struct {
	l1   []int
	l2   []int
	want []int
}

var AddTwoNumbersCases = []AddTwoNumbersCase{
	{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
	{[]int{2, 4, 3}, []int{5, 6, 4, 5}, []int{7, 0, 8, 5}},
	{[]int{}, []int{}, nil},
	{[]int{5}, []int{5}, []int{0, 1}},
	{nil, nil, nil},
}

func makeList(arr []int) *ListNode {
	head := &ListNode{0, nil}
	tail := head

	for _, v := range arr {
		tail.Next = &ListNode{v, nil}
		tail = tail.Next
	}

	return head.Next
}

func convertToSlice(l1 *ListNode) []int {
	var arr []int

	for l1 != nil {
		arr = append(arr, l1.Val)
		l1 = l1.Next
	}

	return arr
}

func TestAddTwoNumbers(t *testing.T) {
	for _, item := range AddTwoNumbersCases {
		result := convertToSlice(addTwoNumbers(makeList(item.l1), makeList(item.l2)))

		if !compareSlice(result, item.want) {
			t.Errorf("Add two numbers result: %v, expect: %v", result, item.want)
		}
	}
}
