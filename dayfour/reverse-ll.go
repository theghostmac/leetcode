package dayfour

type Node struct {
	Value int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList(value int) *LinkedList {
	newNode := &Node{
		Value: value,
	}

	return &LinkedList{
		Head: newNode,
	}
}

func CreateLinkedList(values []int) *LinkedList {
	list := NewLinkedList(values[0])
	current := list.Head

	for _, value := range values[1:] {
		current.Next = NewLinkedList(value).Head
		current = current.Next
	}

	return list
}

func reverseLinkedList(head *Node) *Node {
	var previousNode *Node = nil // pointer to the linkedlist.
	currentNode := head

	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}

	return previousNode
}
