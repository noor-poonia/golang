package main

import "fmt"

// main function => automatically executable
func main()  {
	fmt.Println("Functions in Golang")

	// callinfg the greeter function
	greeter()

	result := adder(3, 5)
	proResult, str := proAdder(3, 4, 5, 6)
	// proResult, _ := proAdder(3, 4, 5, 6)

	fmt.Println("Result is:", result)
	fmt.Println("ProResult is:", proResult)
	fmt.Println("The message is:", str)

	// not allowed to write a function inside a function
	// 	func greeterTwo()  {
	// 		fmt.Println("Another method")
	// 	}
	// 	greeterTwo()
}

func adder(a int, b int) int {
	return a + b
}

// incase you don't know the total count of arguments => all values are of type int
// values is now a slice
// ... => variatic func and can expect any number of arguments
// when returning two values, wrap them in ()
func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Hi Pro result function"
}

func greeter() {
	fmt.Println("Hello from Golang")
}

// anonymous func => it exists with the syntax 
// func () {body}() => executes on the go 

// syntax 
// func <func-name>(value1 dataType, and so on) <dataType-of-expected-return-value> {
	// body
// }