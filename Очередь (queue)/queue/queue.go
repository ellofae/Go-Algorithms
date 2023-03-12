package queue

import "fmt"

// Queue structure
type Node struct {
	Value int
	Next  *Node
}

var size = 0
var Queue *Node = nil

// Get the queue
func GetQueue() *Node {
	return Queue
}

// Get size of the queue
func GetSize() int {
	return size
}

// Add an element to a beginning of the queue
func Push(t *Node, v int) {
	if Queue == nil {
		Queue = &Node{v, nil}
		size++
		return
	}

	t = &Node{v, nil}
	t.Next = Queue
	Queue = t
	size++
}

// Get the first element of the queue and then remove it
func Pop(t *Node) (int, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		Queue = nil
		size--
		return t.Value, true
	}

	temp := t
	for t.Next != nil {
		temp = t
		t = t.Next
	}

	v := temp.Next.Value
	temp.Next = nil

	size--
	return v, true
}

// Print out elements of the queue
func Traverse(t *Node) {
	if size == 0 {
		fmt.Println("Empty Queue!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}
