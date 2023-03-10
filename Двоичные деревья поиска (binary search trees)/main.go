package main

import (
	"bstree"
	"fmt"
)

const filePath = "./tree.txt"

func main() {
	var root *bstree.Tree = nil
	root = bstree.CreateTreeFromFile(root, filePath)

	fmt.Println("root key: ", root.Value)
	bstree.InorderTraverse(root)
	fmt.Println()

	/*
		searchKey := bstree.Search(root, 5)
		if searchKey != nil {
			fmt.Println("key exists: ", searchKey)
		} else {
			fmt.Println("key doesn't exist: ", searchKey)
		}
	*/

	fmt.Println("\ndelete 20: ")
	bstree.DeleteNode(root, 20)
	bstree.InorderTraverse(root)
	fmt.Println()

	fmt.Println("\ndelete 30: ")
	bstree.DeleteNode(root, 30)
	bstree.InorderTraverse(root)
	fmt.Println()

	fmt.Println("\ndelete 50: ")
	bstree.DeleteNode(root, 50)
	bstree.InorderTraverse(root)
	fmt.Println()

	sliceInt := make([]int, 0)
	longPath := bstree.LongestPath(root, sliceInt)

	fmt.Println("temp: ", longPath)
}
