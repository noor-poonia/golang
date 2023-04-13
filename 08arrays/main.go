package main

import "fmt"

func main()  {

	// array => very less used in golang
	// under the hood other ds use arrays 
	fmt.Println("Arrays")

	// syntax => always specify the length of the array
	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[3] = "peach"
	// since we didn't specify fruitList[2], the output has one space left
	fmt.Println("Fruit list is", fruitList)
	// length will be 4 since we reserved that memory even tho we have only 3 elts in the array
	fmt.Println("Fruit list is", len(fruitList))

	// antoehr syntax
	var vegList = [3]string{"potato", "beans", "mushrooms"}
	fmt.Println("Veg List is", vegList)
	fmt.Println("Veg List is", len(vegList))

}