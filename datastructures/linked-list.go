package datastructures

import "fmt"

// Node represents a single node in the linked list.
type Node struct {
	Data     int
	NextNode *Node
}

// Inserting adds a new node with the given data to the list at a specific position.
func (head *Node) Inserting(data int, position int) {
	newNode := &Node{
		Data: data,
	}

	// handle insertion at the beginning.
	if position == 0 {
		newNode.NextNode = head
		head = newNode
		return
	}

	// traverse to the node before the insertion position.
	current := head
	for i := 1; i < position && current != nil; i++ {
		current = current.NextNode
	}

	// insert the new node.
	if current != nil {
		newNode.NextNode = current.NextNode
		current.NextNode = newNode
	} else {
		fmt.Println("invalid position for insertion.")
	}
}

// Deleting removes an existing node from the list.
func (head *Node) Deleting() {

}

// Searching finds a specific element in the list.
func (head *Node) Searching() {

}

// Traversing visits each node in the list in a sequential order.
func (head *Node) Traversing() {

}

// Sorting rearranges the node in a specific order
func (head *Node) Sorting() {

}

// Reversing reverses the order of the nodes in the list.
func (head *Node) Reversing() {

}
