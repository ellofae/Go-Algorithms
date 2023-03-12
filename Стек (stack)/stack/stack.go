package stack

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var size = 0
var Stack *Node = nil

func GetSize() int {
	return size
}

func Push(v int) {
	if Stack == nil {
		Stack = &Node{v, nil}
		size = 1
		return
	}

	newNode := &Node{v, nil}
	newNode.Next = Stack
	Stack = newNode
	size++

	return
}

func Pop(t *Node) (int, bool) {
	if t == nil {
		return 0, false
	}

	if size == 1 {
		size = 0
		Stack = nil
		return t.Value, true
	}

	Stack = Stack.Next
	size--
	return t.Value, true
}

func PrintStack(t *Node) {
	if t == nil {
		fmt.Println("The stack is empty!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}

	fmt.Println()
}
