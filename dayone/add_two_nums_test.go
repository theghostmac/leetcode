package dayone

import "testing"

func Test_AddTwoNumbers(t *testing.T) {
	cases := []struct {
		l1, l2, expected *ListNode
	}{
		// Basic cases
		{createList(2, 4, 3), createList(5, 6, 4), createList(7, 0, 8)},
		{createList(0), createList(0), createList(0)},

		// Lists of different lengths
		{createList(9, 9, 9), createList(1), createList(0, 0, 0, 1)},
		{createList(1), createList(9, 9), createList(0, 0, 1)},

		// Carry-overs
		{createList(1, 8), createList(0), createList(1, 8)},
		{createList(9), createList(1), createList(0, 1)},

		// Empty lists
		{nil, createList(0), createList(0)},
		{createList(0), nil, createList(0)},
		{nil, nil, nil},
	}

	for _, tc := range cases {
		result := addTwoNumbers(tc.l1, tc.l2)
		if !equalLists(result, tc.expected) {
			t.Errorf("addTwoNumbers(%v, %v) = %v, expected %v", tc.l1, tc.l2, result, tc.expected)
		}
	}
}

func createList(values ...int) *ListNode {
	var head *ListNode
	tail := &head
	for _, value := range values {
		*tail = &ListNode{Val: value}
		tail = &(*tail).Next
	}

	return head
}

func equalLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	return l1 == nil && l2 == nil
}
