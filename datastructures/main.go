package main

import "fmt"

func main() {
	head := CreateNode(10)
	head.Inserting(20, 1)
	head.Inserting(30, 2)

	head.Traversing()

	nodeWith20 := head.Searching(20)
	fmt.Println(nodeWith20)
}
