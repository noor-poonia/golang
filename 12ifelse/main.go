package main 

import "fmt"

func main()  {
    fmt.Println("If-else in Golang")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10{
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println("Result:", result)

	// special syntax => most used in web handling 
	// syntax is basically initializing it and assigning the value at the same time
	// this is separated by ; after which we checking the if condition on the go 
	if num := 3; num < 10 {
		fmt.Println("num is less 10")
	} else {
		fmt.Println("num is not less than 10")
	}
}