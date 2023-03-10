package btree

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// const filePath = `./tmp/btree.txt`
const filePath = `./tmp/btree.txt`

var (
	maxIncCounter = 0
	maxDecCounter = 0
)

type Node struct {
	Key   int
	Left  *Node
	Right *Node
	Prev  *Node
}

func createNode(key int) *Node {
	return &Node{
		Key:   key,
		Left:  nil,
		Right: nil,
		Prev:  nil,
	}
}

func insertLeft(root *Node, elem *Node) {
	if root.Left == nil {
		elem.Prev = root
		root.Left = elem
	}
}

func insertRight(root *Node, elem *Node) {
	if root.Right == nil {
		elem.Prev = root
		root.Right = elem
	}
}

func goToLeft(root *Node) *Node {
	if root.Left != nil {
		return root.Left
	} else {
		return root
	}
}

func goToRight(root *Node) *Node {
	if root.Right != nil {
		return root.Right
	} else {
		return root
	}
}

func goPrev(root *Node) *Node {
	if root.Prev != nil {
		return root.Prev
	} else {
		return root
	}
}

/*
Commands:
cl - create left
cr - create right
gl - go left
gr - go right
gp - go previous
*/
func CreateFromFile(root *Node) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Didn't manage to open the file.")
		os.Exit(1)
	}
	defer f.Close()

	counter := 0
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error during reading the file.")
			os.Exit(1)
		}

		splitedLine := strings.Fields(line)
		convValue, err := strconv.Atoi(splitedLine[0])
		if err != nil {
			fmt.Println("Error during parsing the value in line")
			os.Exit(1)
		}

		if counter == 0 {
			root.Key = convValue
			counter++
			continue
		}

		switch splitedLine[1] {
		case "cl":
			insertLeft(root, createNode(convValue))
		case "cr":
			insertRight(root, createNode(convValue))
		case "gl":
			root = goToLeft(root)
		case "gr":
			root = goToRight(root)
		case "gp":
			root = goPrev(root)
		}
	}
}

func PrintTree(root *Node, level int) {
	if root == nil {
		return
	} else {
		level++
		PrintTree(root.Right, level)

		for i := 0; i < level; i++ {
			fmt.Print("|")
		}

		fmt.Println(root.Key)

		PrintTree(root.Left, level)

	}
}

func longestPathIncrease(root *Node, expendingSlice []int, mainSlice [][]int) [][]int {
	if len(expendingSlice) == 0 {
		expendingSlice = append(expendingSlice, root.Key)
	} else {
		if expendingSlice[len(expendingSlice)-1] <= root.Key {
			expendingSlice = append(expendingSlice, root.Key)
		} else {
			if !checkSliceExists(mainSlice, expendingSlice) {
				mainSlice = append(mainSlice, expendingSlice)
			}

			expendingSlice = make([]int, 0)
			expendingSlice = append(expendingSlice, root.Key)
			//	}
			//fmt.Println("test: ", !checkSliceExists(mainSlice, expendingSlice), "main: ", mainSlice, "sub: ", expendingSlice)
		}
	}

	if len(expendingSlice) > maxIncCounter {
		maxIncCounter = len(expendingSlice)
	}

	if root.Left != nil {
		mainSlice = longestPathIncrease(root.Left, expendingSlice, mainSlice)
	}

	if root.Right != nil {
		mainSlice = longestPathIncrease(root.Right, expendingSlice, mainSlice)
	}

	if root.Left == nil && root.Right == nil {
		mainSlice = append(mainSlice, expendingSlice)
	}
	return mainSlice
}

func longestPathDecrease(root *Node, expendingSlice []int, mainSlice [][]int) [][]int {
	if len(expendingSlice) == 0 {
		expendingSlice = append(expendingSlice, root.Key)
	} else {
		if expendingSlice[len(expendingSlice)-1] >= root.Key {
			expendingSlice = append(expendingSlice, root.Key)
		} else {
			if !checkSliceExists(mainSlice, expendingSlice) {
				mainSlice = append(mainSlice, expendingSlice)
			}

			expendingSlice = make([]int, 0)
			expendingSlice = append(expendingSlice, root.Key)

		}
	}

	if len(expendingSlice) > maxDecCounter {
		maxDecCounter = len(expendingSlice)
	}

	if root.Left != nil {
		mainSlice = longestPathDecrease(root.Left, expendingSlice, mainSlice)
	}

	if root.Right != nil {
		mainSlice = longestPathDecrease(root.Right, expendingSlice, mainSlice)
	}

	if root.Left == nil && root.Right == nil {
		mainSlice = append(mainSlice, expendingSlice)
	}

	return mainSlice
}

func LongestPath(root *Node) {
	expedingSlice := make([]int, 0)
	mainSlice := make([][]int, 0)

	tempIncrease := longestPathIncrease(root, expedingSlice, mainSlice)
	tempDecrease := longestPathDecrease(root, expedingSlice, mainSlice)

	fmt.Println("\ntempInc: ", tempIncrease)
	fmt.Println("tempDec: ", tempDecrease)

	fmt.Println()

	paths := pathConcatination(tempIncrease, tempDecrease)
	getLongestPaths(paths)

	//fmt.Println("paths: ", paths)

}

func pathConcatination(inc, dec [][]int) [][]int {
	path := make([][]int, 0)

	for i := 0; i < len(dec); i++ {
		for j := 0; j < len(inc); j++ {
			if dec[i][0] == inc[j][0] {
				tempSlice := make([]int, len(dec[i])+len(inc[j])-1)
				tempCount := 0
				for j := len(dec[i]) - 1; j >= 0; j-- {
					tempSlice[tempCount] = dec[i][j]
					tempCount++
				}

				tempSlice = append(tempSlice[:tempCount-1], inc[j]...)

				if !checkSliceExists(path, tempSlice) {
					path = append(path, tempSlice)
				}

				tempSlice = nil
			} else {
				if len(dec[i]) > len(inc[j]) {
					tempSlice := make([]int, len(dec[i]))
					tempSlice = append(tempSlice[:0], dec[i]...)

					if !checkSliceExists(path, tempSlice) {
						path = append(path, tempSlice)
					}

					tempSlice = nil
				} else {
					tempSlice := make([]int, len(inc[j]))
					tempSlice = append(tempSlice[:0], inc[j]...)

					if !checkSliceExists(path, tempSlice) {
						path = append(path, tempSlice)
					}

					tempSlice = nil
				}
			}

		}
	}
	return path
}

func getLongestPaths(paths [][]int) {
	maxLen := 0
	ind := 0

	for _, value := range paths {
		if len(value) > maxLen {
			maxLen = len(value)
		}
	}

	for _, value := range paths {
		if len(value) == maxLen {
			fmt.Printf("path %d: %v\n", ind, value)
			ind++
		}
	}
}

func checkSliceExists(mainSlice [][]int, slice []int) bool {
	for _, subslice := range mainSlice {
		if len(subslice) == len(slice) {
			check := true
			for j, elem := range subslice {
				if elem != slice[j] {
					check = false
				}
			}

			if check {
				return true
			}
		}
	}

	return false
}
