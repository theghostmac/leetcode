package main

import "fmt"

// Node represents a single node in the linked list.
type Node struct {
	Data     int
	NextNode *Node
}

func CreateNode(data int) *Node {
	return &Node{Data: data}
}

// Inserting adds a new node with the given data to the list at a specific position.
func (head *Node) Inserting(data int, position int) {
	// create a new node with that data.
	newNode := &Node{
		Data: data,
	}

	// if there is no head (position = 0), make the new node the head node.
	// update: i guess this rewrites the head, because of the new CreateNode function above.
	if position == 0 {
		newNode.NextNode = head
		head = newNode
		return
	}

	// move to the specified position.
	current := head
	for i := 1; i < position && current != nil; i++ {
		current = current.NextNode
	}

	// insert the node if the position exists.
	if current != nil {
		newNode.NextNode = current.NextNode
		current.NextNode = newNode
	} else /*  else return info that node doesn't exist. */ {
		fmt.Println("invalid position for insertion")
	}
}

// Deleting removes an existing node from the list.
func (head *Node) Deleting(position int) {

}

// Searching finds a specific element in the list.
func (head *Node) Searching(data int) *Node {
	current := head
	for current != nil {
		if current.Data == data {
			return current
		}

		current = current.NextNode
	}

	return nil // not found.
}

// Traversing visits each node in the list in a sequential order.
func (head *Node) Traversing() {
	current := head
	for current != nil {
		fmt.Println(current.Data)
		current = current.NextNode
	}
}

// Sorting rearranges the node in a specific order
func (head *Node) Sorting() {

}

// Reversing reverses the order of the nodes in the list.
func (head *Node) Reversing() {

}
