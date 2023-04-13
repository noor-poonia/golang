package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our app")
	fmt.Println("Please rate our service between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)

	// strconv is a package used for type conversion 
	// ParseFloat takes 2 values => 1st the value to be modified and 2nd the bit size which is usually is 64 in float
	// it returns 2 things so we need to use comma ok or error syntax
	// strings is another package that is used to trim the extra \n that is being added to the input
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// explicit type conversion
	// var geek2 float64 = float64(geek1)
	// The general syntax for converting a value val to a type T is T(val). 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numRating + 1)
	}

	// GO lang is a strong typed language and so it doesn't support implicit type conversion

	fmt.Printf("Type of error is %T \n", err)

	num := 10

	var number float32 = float32(num)

	fmt.Printf("Type of number is %T \n", number)
}