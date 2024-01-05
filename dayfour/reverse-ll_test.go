package dayfour

import (
	"testing"
)

func Test_ReverseLinkedList(t *testing.T) {
    testCases := []struct {
        input    *Node
        expected *Node
    }{
        // Empty list
        {nil, nil},
        // Single node
        {NewLinkedList(1).Head, NewLinkedList(1).Head},
        // Multiple nodes
        {CreateLinkedList([]int{1, 2, 3, 4, 5}).Head, CreateLinkedList([]int{5, 4, 3, 2, 1}).Head},
    }

    for _, tc := range testCases {
        actual := reverseLinkedList(tc.input)
        if !checkLinkedListsEqual(actual, tc.expected) {
            t.Errorf("Reverse LinkedList (%v) = %v, expected %v", tc.input, actual, tc.expected)
        }
    }
}

// Helper function to compare linked lists
func checkLinkedListsEqual(list1, list2 *Node) bool {
    for list1 != nil && list2 != nil {
        if list1.Value != list2.Value {
            return false
        }
        list1 = list1.Next
        list2 = list2.Next
    }
    return list1 == nil && list2 == nil
}
