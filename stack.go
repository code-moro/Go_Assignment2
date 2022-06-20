package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	n   int64
	top int64
	arr [100]int64
}

// func to pop the values  from stack
func pop(s *Stack) int64 {
	var r int64
	if s.top <= -1 {
		fmt.Println("Stack underFlow")
		return -1
	} else {
		r = s.arr[s.top]
		s.arr[s.top] = 0
		s.top--
	}
	return r
}
// func to push values in stack
func push(s *Stack, val int64) {
	if s.top >= s.n-1 {
		fmt.Println("Stack overflow")
	} else {
		s.top++
		s.arr[s.top] = val
	}
}
// func to Display values from stack
func Display(s *Stack) {
	var i int64
	for i = 0; i <= s.top; i++ {
		fmt.Print(s.arr[i], "\t")
	}
}


func main() {
	reader := bufio.NewScanner(os.Stdin)
	S := Stack{}
	fmt.Print("Enter the Size Of Stack :")
	reader.Scan()
	nval, _ := strconv.ParseInt(reader.Text(), 10, 64)
	S.n = nval
	S.top = -1
	fmt.Println("1.Push Element onto the stack\n2.Pop Elemenet form the stack\n3.Display Elements in stack\n4.Exit Program")
	for {
		fmt.Print("Enter the choice :")
		reader.Scan()
		choice, _ := strconv.ParseInt(reader.Text(), 10, 64)

		switch choice {

		case 1:
			fmt.Print("Enter the Number you want to add:")
			reader.Scan()
			val, _ := strconv.ParseInt(reader.Text(), 10, 64)
			push(&S, val)

		case 2:
			ele := pop(&S)
			if ele > -1 {
				fmt.Println("Deleted Element From Stack is :", ele)
			}

		case 3:
			fmt.Println("Stack Elements:")
			Display(&S)
			fmt.Println()

		case 4:
			fmt.Println("Program Exited!!")
			return

		default:
			fmt.Println("Invalid choice!!")
		}

	}

}
