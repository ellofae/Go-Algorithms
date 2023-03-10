package bstree

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func Test() {
	fmt.Println("test")
}

type Tree struct {
	Left  *Tree
	Right *Tree
	Value int
}

// Insert a node in the binary search tree
func Insert(node *Tree, n int) *Tree {
	if node == nil {
		return &Tree{Left: nil, Right: nil, Value: n}
	}
	if n == node.Value {
		return node
	}
	if n < node.Value {
		node.Left = Insert(node.Left, n)
		return node
	}
	node.Right = Insert(node.Right, n)
	return node
}

// Creation of a binary search tree via the random numbers
func CreateTreeRand(n int) *Tree {
	var root *Tree
	for i := 0; i < 2*n; i++ {
		randValue := rand.Intn(2 * n)
		root = Insert(root, randValue)
	}
	return root
}

// Creation of a binary search tree via the file reading
func CreateTreeFromFile(root *Tree, filename string) *Tree {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Didn't manage to open the file")
		os.Exit(1)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error occured during reading the file.")
			os.Exit(1)
		}

		value, err := strconv.Atoi(line[:len(line)-1])
		if err != nil {
			log.Println("file contains a non-valid character for the tree: ", line)
			continue
		}
		root = Insert(root, value)
	}

	return root
}

// Printing out the nodes from the left leaf to the right leaf
func InorderTraverse(node *Tree) {
	if node == nil {
		return
	}
	InorderTraverse(node.Left)
	fmt.Printf("%d ", node.Value)
	InorderTraverse(node.Right)
}

// Searching for a key in the binary search tree
func Search(node *Tree, key int) *Tree {
	if node == nil || node.Value == key {
		return node
	}

	if node.Value < key {
		return Search(node.Right, key)
	} else {
		return Search(node.Left, key)
	}
}

// Delete a node with the specific key
func DeleteNode(root *Tree, key int) *Tree {
	if root == nil {
		return root
	}

	if root.Value < key {
		root.Right = DeleteNode(root.Right, key)
	} else if root.Value > key {
		root.Left = DeleteNode(root.Left, key)
	} else {
		if root.Left == nil {
			root := root.Right
			return root
		} else if root.Right == nil {
			root = root.Left
			return root
		}

		temp := MinValueNode(root.Right)
		root.Value = temp.Value

		root.Right = DeleteNode(root.Right, root.Value)
	}
	return root
}

// Get the node with the most minimum key
func MinValueNode(node *Tree) *Tree {
	if node.Left == nil {
		return node
	}

	return MinValueNode(node.Left)
}

func longestPathUtilLeft(root *Tree, longestPath []int) []int {
	if root == nil {
		return longestPath
	}

	if longestPath[len(longestPath)-1] >= root.Value {
		longestPath = append(longestPath, root.Value) // recived a new address

		rightPath := longestPathUtilLeft(root.Right, longestPath)
		leftPath := longestPathUtilLeft(root.Left, longestPath)

		//fmt.Println("leftPath: ", leftPath)
		//fmt.Println("rightPath: ", rightPath)

		if len(leftPath) > len(rightPath) {
			return leftPath
		} else {
			return rightPath
		}
	} else {
		return nil
	}
}

func longestPathUtilRight(root *Tree, longestPath []int) []int {
	if root == nil {
		return longestPath
	}

	if longestPath[len(longestPath)-1] <= root.Value {
		longestPath = append(longestPath, root.Value) // recived a new address

		rightPath := longestPathUtilRight(root.Right, longestPath)
		leftPath := longestPathUtilRight(root.Left, longestPath)

		//fmt.Println("leftPath: ", leftPath)
		//fmt.Println("rightPath: ", rightPath)

		if len(leftPath) > len(rightPath) {
			return leftPath
		} else {
			return rightPath
		}
	} else {
		return nil
	}
}

// Print out the longest path (paths) from the leafs to nodes or nodes to nodes in increasing order
func LongestPath(root *Tree, longestPath []int) []int {
	if len(longestPath) == 0 {
		longestPath = append(longestPath, root.Value) // adding the root
		return LongestPath(root, longestPath)
	} else if root == nil {
		return longestPath
	}

	leftPath := longestPathUtilLeft(root.Left, longestPath)
	rightPath := longestPathUtilRight(root.Right, longestPath)

	fmt.Println("left path: ", leftPath)
	fmt.Println("left path: ", rightPath)

	// reverse the left path
	for i, j := 0, len(leftPath)-1; i < len(leftPath)/2; i, j = i+1, j-1 {
		leftPath[i], leftPath[j] = leftPath[j], leftPath[i]
	}

	// concatinate the paths
	leftPath = leftPath[:len(leftPath)-1]
	fullPath := make([]int, len(leftPath)+len(rightPath))
	for i := 0; i < len(leftPath); i++ {
		fullPath[i] = leftPath[i]
	}

	for i := 0; i < len(rightPath); i++ {
		fullPath[i+len(leftPath)] = rightPath[i]
	}

	//fmt.Println("fullPath: ", fullPath)

	return fullPath
}
