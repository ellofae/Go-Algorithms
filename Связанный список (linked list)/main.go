package main

import (
	"fmt"
	"linked"
)

func main() {
	root := linked.GetHeadPtr()

	for i := 0; i < 10; i++ {
		root = linked.Enqueue(root, i*10)
	}
	linked.Traverse(root)

	ok := linked.LookupNode(root, 100)
	fmt.Printf("%d exists?: %v\n", 100, ok)

	ok = linked.LookupNode(root, 60)
	fmt.Printf("%d exists?: %v\n", 60, ok)

	fmt.Printf("Size: %d\n\n", linked.GetListSize(root))

	fmt.Println("Removing the last node:")
	root = linked.Dequeue(root)
	linked.Traverse(root)
	fmt.Printf("Size: %d\n", linked.GetListSize(root))

	root = linked.Dequeue(root)
	linked.Traverse(root)
	fmt.Printf("Size: %d\n", linked.GetListSize(root))

	fmt.Println("Add 25: ")
	root = linked.InsertNode(root, &linked.Node{25, nil})
	linked.Traverse(root)

	fmt.Println("Add 27: ")
	root = linked.InsertNode(root, &linked.Node{27, nil})
	linked.Traverse(root)

	fmt.Println("Add -10: ")
	root = linked.InsertNode(root, &linked.Node{-10, nil})
	linked.Traverse(root)

	fmt.Println("Add -20: ")
	root = linked.InsertNode(root, &linked.Node{-20, nil})
	linked.Traverse(root)

}
