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

func (ll *LinkedList) DeleteNode() {}

func (ll *LinkedList) SearchNode() {}
