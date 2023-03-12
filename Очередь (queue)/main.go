package main

import (
	"fmt"
	q "queue"
)

func main() {
	for i := 1; i <= 10; i++ {
		q.Push(q.Queue, i*10)
	}
	q.Traverse(q.Queue)
	fmt.Printf("size: %d\n", q.GetSize())

	value, ident := q.Pop(q.Queue)
	if ident {
		fmt.Println("\nPop:", value)
	}

	value, ident = q.Pop(q.Queue)
	if ident {
		fmt.Println("Pop:", value)
	}

	fmt.Println("\nSize: ", q.GetSize())
	q.Traverse(q.Queue)
}
