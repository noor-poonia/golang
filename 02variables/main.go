package main

import "fmt"

// this is not allowed bcz this operator is allowed inside methods but not outside
// jwtToken := 30000

// since the variable name is starting with a capital letter, its almost equal to writing public before variable
// its now accessible anywhere in this file or this folder
const LginToken string = "dljfheskfhj"

func main() {
	var username string = "noor"
	fmt.Println(username);
	fmt.Printf("Variable is of type: %T \n", username)

	var isLggedIn bool = false
	fmt.Println(isLggedIn);
	fmt.Printf("Variable is of type: %T \n", isLggedIn)

	// uint8 boundary line is 0 to 255 -> considering binary representation
	// unit16, unit32 also follow  the same pattern
	// int data type mostly covers all these values
	var smallVal uint8 = 255
	fmt.Println(smallVal);
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// float32 and float64 have precision difference
	// float64 is more precise than float32
	var smallFloat float32 = 255.4554452334566565435
	fmt.Println(smallFloat);
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and aliases
	// default values are never garbage values in go
	// for all data types 
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// another way of declaring variables => implicit type => lexer decides it according to the value passed
	var website = "www.google.com"
	fmt.Println(website)

	// another way => no var style => most common style
	// := is known as the walrus operator
	numberOfUsers := 3000
	fmt.Println(numberOfUsers)

	fmt.Println(LginToken)
}