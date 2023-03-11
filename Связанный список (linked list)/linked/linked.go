package linked

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

// The head of the linked list
var root *Node = nil

// Get the head of the linked list
func GetHeadPtr() *Node {
	return root
}

// Add new node at the end of the list
func Enqueue(t *Node, v int) *Node {
	if t == nil {
		root = &Node{Value: v, Next: nil}
		return root
	}

	if t.Value == v {
		fmt.Printf("Node already exists: %d\n", v)
		return root
	}

	if t.Next == nil {
		t.Next = &Node{Value: v, Next: nil}
		return root
	}

	return Enqueue(t.Next, v)
}

// Removes the last element in the queue
func Dequeue(root *Node) *Node {
	if root == nil {
		fmt.Println("Empty list!")
		return nil
	}

	tempHead := root
	var prevNode *Node

	for root.Next != nil {
		prevNode = root
		root = root.Next
	}
	prevNode.Next = nil

	return tempHead
}

// Insert a node
func InsertNode(root *Node, t *Node) *Node {
	if root == nil {
		fmt.Println("Empty list!")
		return nil
	}

	if root.Value > t.Value {
		temp := root

		root = t
		root.Next = temp

		return root
	}

	tempHead := root

	for root.Next != nil {
		if root.Value < t.Value && t.Value < root.Next.Value {
			temp := root.Next

			root.Next = t
			root.Next.Next = temp

			return tempHead
		} else {
			root = root.Next
		}
	}

	return tempHead

}

// Check if the linked list is empty or not
func IsEmpty(root *Node) bool {
	if root == nil {
		return true
	}
	return false
}

// Print out the nodes of the list
func Traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list")
		return
	}

	for t != nil {
		fmt.Printf("-> %d ", t.Value)
		t = t.Next
	}

	fmt.Println()
}

// Check if the node exists in the list
func LookupNode(t *Node, v int) bool {
	if t == nil {
		t = &Node{Value: v, Next: nil}
		root = t
		return false
	}

	if t.Value == v {
		return true
	}

	if t.Next == nil {
		return false
	}

	return LookupNode(t.Next, v)
}

// Get size of the linked list
func GetListSize(t *Node) int {
	if t == nil {
		fmt.Println("Empty list!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}

	return i
}
