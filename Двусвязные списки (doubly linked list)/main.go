package main

import (
	"dlinked"
	"fmt"
)

func main() {
	root := dlinked.GetHeadPtr()

	for i := 1; i < 10; i++ {
		root = dlinked.Enqueue(root, i*10)
	}

	dlinked.Traverse(root)
	fmt.Println("size: ", dlinked.GetListSize(root))

	fmt.Println("\nRemoving last element:")
	root = dlinked.Dequeue(root)
	dlinked.Traverse(root)

	root = dlinked.Dequeue(root)
	dlinked.Traverse(root)

	fmt.Println("\n Reversed:")
	dlinked.Reverse(root)
}
