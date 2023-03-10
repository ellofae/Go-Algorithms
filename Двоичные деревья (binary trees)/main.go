package main

import (
	"btree"
)

func main() {
	root := new(btree.Node)
	btree.CreateFromFile(root)
	btree.PrintTree(root, 0)

	btree.LongestPath(root)

}
