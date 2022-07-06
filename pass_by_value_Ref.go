package main

import "fmt"

type values struct {
	x int
	y int
}

func ByValue(v values, a int, b int) {

	v.x = a
	v.y = b

}

func ByRef(v *values, a int, b int) {

	v.x = a
	v.y = b

}

func main() {
	v := values{
		x: 10,
		y: 20,
	}
	fmt.Println("Original values of x", v.x,"value of y", v.y)
	ByValue(v, 1, 5)
	fmt.Println("After call by value func values of x", v.x,"value of y", v.y)
	ByRef(&v, 30, 40)
	fmt.Println("After call by ref func values of  x", v.x,"value of y", v.y)
}
