package main

import "fmt"

func main()  {

	// the flow of execution of a function is line by line 
	// but when we put the keyword "defer", its not going to execute in that flow but at the very end
	// its like LIFO structure
	fmt.Println("Defer in golang")

	// executed at the end
	// according to the lifo structure, the 2nd defer stmt is executed 1st and then the 1st defer stmt
	defer fmt.Println("1st defer")
	defer fmt.Println("2nd defer")

	fmt.Println("Hello Golang")

	deferFunc()
	fmt.Println(add())
}

func deferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func add() string {
	defer func () string {
		return "hi"
	}()
	return "hello"
}