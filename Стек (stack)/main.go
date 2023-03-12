package main

import (
	"fmt"
	s "stack"
)

func main() {
	for i := 1; i <= 10; i++ {
		s.Push(i * 10)
	}

	s.PrintStack(s.Stack)
	fmt.Println("Size: ", s.GetSize())

	value, ident := s.Pop(s.Stack)
	if ident {
		fmt.Println("Pop: ", value)
	} else {
		fmt.Println("Pop() failed!")
	}

	value, ident = s.Pop(s.Stack)
	if ident {
		fmt.Println("Pop: ", value)
	} else {
		fmt.Println("Pop() failed!")
	}
	fmt.Println()

	s.PrintStack(s.Stack)
	fmt.Println("Size: ", s.GetSize())
}
