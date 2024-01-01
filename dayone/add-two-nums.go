package dayone

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// create an empty node to store the results list.
	dummyHead := &ListNode{}
	// create a pointer to the empty node, used to traverse the result list.
	currentNode := dummyHead
	// initialize variable to track carry-overs from addition.
	carry := 0

	// loop to iterate over each list node.
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		currentNode.Next = &ListNode{Val: sum % 10}
		currentNode = currentNode.Next
	}

	return dummyHead.Next
}
