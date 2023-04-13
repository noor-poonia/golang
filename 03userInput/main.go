package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// bufio => package for I/O

	welcome := "Welcome to user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age: ")

	// comma ok syntax || error ok syntax => bcz there is no try catch syntax in go 
	// keep reading until there is a new line
	// this _ denotes that we only care abt the input value and anything else that comes along can be ignored
	// if u care abt the other things then instead of _ u can use a variable name
	// _ can also be used in place of input depending on the situation
	input, _ := reader.ReadString('\n') 
	fmt.Println("Age is", input)
	fmt.Printf("Type of age is %T \n", input)

	// incase u want to conside both the error and input
	// if there is no error, you will get <nil> as output 
	// comparing to try catch => input works like try and err/_ works like catch
	input, err := reader.ReadString('\n') 
	fmt.Println(err)
	fmt.Println("Age is", input)
	fmt.Printf("Type of age is %T \n", input)
}