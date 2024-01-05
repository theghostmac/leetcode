package linkedlist

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList(value int) *LinkedList {
	headNode := &Node{
		Value: value,
	}
	return &LinkedList{
		Head: headNode,
	}
}

func (ll *LinkedList) String() string {
	var str string
	current := ll.Head
	for current != nil {
		str += fmt.Sprintf("%v -> ", current.Value)
		current = current.Next
	}

	str += "nil"

	return str
}

func (ll *LinkedList) InsertNewNode(value int) {
	newNode := &Node{
		Value: value,
		Next:  ll.Head,
	}

	ll.Head = newNode
}

func (ll *LinkedList) InsertNodeAtPosition(value int, position int) {
	if position < 0 {
		fmt.Println("invalid position: position cannot be negative")
		return
	}

	if position == 0 {
		ll.InsertNewNode(value)
		return
	}

	previousNode, currentNode := ll.TraverseToX(position - 1)

	if currentNode == nil {
		fmt.Println("Invalid position: Position exceeds list length.")
		return
	}

	newNode := &Node{
		Value: value,
		Next:  currentNode,
	}

	previousNode.Next = newNode
}

func (ll *LinkedList) TraverseToX(x int) (*Node, *Node) {
	if x < 0 {
		fmt.Println("invalid operation: p")
		return nil, nil
	}

	var previousNode *Node = nil
	currentNode := ll.Head

	for i := 1; i <= x && currentNode != nil; i++ {
		previousNode = currentNode
		currentNode = currentNode.Next
	}

	return previousNode, currentNode
}

func (ll *LinkedList) DeleteDuplicates(node *Node) *Node {
	currentNode := node

	for currentNode != nil && currentNode.Next != nil {
		runner := currentNode.Next
		for runner != nil && runner.Value == currentNode.Value {
			runner = runner.Next
		}

		currentNode.Next = runner
		currentNode = currentNode.Next
	}

	return node
}

func (ll *LinkedList) MergeTwoSortedLists(l1, l2 *LinkedList) *LinkedList {
	// create a new empty linked list.
	mergedList := NewLinkedList(0)
	currentNode := mergedList.Head

	// initialize two pointers to the head of both l1 & l2.
	p1, p2 := l1.Head, l2.Head

	// iterate through both lists simulatanously.
	for p1 != nil && p2 != nil {
		// if p1's value is smaller, add to the merged list, then advance to the next node.
		if p1.Value < p2.Value {
			currentNode.Next = p1
			p1 = p1.Next
		} else {
			currentNode.Next = p2
			p2 = p2.Next
		}

		currentNode = currentNode.Next
	}

	// if either pointer reaches the end of its list, add the remaining nodes of the second list to the merged list.
	if p1 != nil {
		currentNode.Next = p1
	} else {
		currentNode.Next = p2
	}

	// return the merged list.
	return mergedList
}

func (ll *LinkedList) DeleteNode() {}

func (ll *LinkedList) SearchNode() {}

func (ll *LinkedList) ReverseALinkedList(linkedList *LinkedList) {
	// reversedList := NewLinkedList(0) // create empty linkedlist.
	// currentNode := reversedList.Head
}
