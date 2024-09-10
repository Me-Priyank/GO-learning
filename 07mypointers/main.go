package main

import "fmt"

// Pointer gives memory referance of variable instead of variable itself, so as to avoind irregularities . because at many places , a copy of that variable is passed instead of the original one.

func main() {
	fmt.Println("Welcome to pointers")

	var ptr *int 
	fmt.Println("value is ",ptr)
}
