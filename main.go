package main

import (
	"fmt"

	"github.com/theghostmac/leetcode/ds"
)

func main() {
	newLL := ds.NewLinkedList(10)
	fmt.Println(newLL.String())

	newLL.InsertNewNode(20)
	fmt.Println(newLL.String())
}
