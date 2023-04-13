package main

import "fmt"

func main()  {

	// pointers give us the assurity that its passing the memory address of these variables (actual value from memory)
	// its the direct reference to the memory address
	fmt.Println("My Pointers")

	// ptr is a pointer responsible for holding integers in it
	var ptr *int
	// * is used to create a pointer
	// the output is <nil> if nthg is initialized
	fmt.Println("Value of pointer is", ptr)

	myNumber := 23
	// & is also used to create a pointer by referring it to a value/memory
	// reference => &
	var ptr1 = &myNumber
	// this gives the actual memory address
	fmt.Println("Value of actual pointer is", ptr1)
	// this gives the value inside the pointer
	fmt.Println("Value of actual pointer is", *ptr1)

	*ptr1 = *ptr1 + 2
	fmt.Println("Result is", myNumber)

	//everything in go is pass by value => what you see is what you get

}