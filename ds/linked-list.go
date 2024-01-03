package ds

import "fmt"

type Node struct {
	Value int
	Next *Node
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
		Next: ll.Head,
	}
	
	ll.Head = newNode
}

func (ll *LinkedList) InsertNodeAtPosition(value int, position int) {
	ll.TraverseToX(position)

	newNode := &Node{
		Value: value,
	}
}

func (ll *LinkedList) TraverseToX(x int) (*Node, *Node) {
	if x < 0 {
		fmt.Println("Invalid position: position cannot be negative")
		return nil, nil
	}

	var prevNode *Node = nil
	currentNode := ll.Head
	for i := 1; i <= x && currentNode != nil; i++ {
		prevNode = currentNode
		currentNode = currentNode.Next
	}

	return prevNode, currentNode
}

func (ll *LinkedList) DeleteNode() {}

func (ll *LinkedList) SearchNode() {}
