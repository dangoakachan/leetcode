package problems

// Problem 2. Add Two Numbers
// URL: https://leetcode.com/problems/add-two-numbers/tabs/description

// ListNode defines a singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil} // sentinel node
	carry := 0                // carray bit

	for tail := head; l1 != nil || l2 != nil || carry != 0; tail = tail.Next {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		tail.Next, carry = &ListNode{sum % 10, nil}, sum/10
	}

	return head.Next // discard sentinel node
}
