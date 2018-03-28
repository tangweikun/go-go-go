package main

import "fmt"

func foo() (int, int) {
	return 3, 7
}

func main() {
	a, b := foo()
	fmt.Println(a, b)

	_, c := foo()
	fmt.Println(c)
}