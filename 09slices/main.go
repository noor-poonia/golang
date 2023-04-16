package main

import (
	"fmt"
	"sort"
)

func main()  {

	// slices => most commonly used data structure
	// under the hood they are arrays => these arrays get some abstraction layers and then they are called slices 
	// slices are more powerful than arrays
	// slices make it easy to manage databases
	fmt.Println("Slices in Golang")

	// syntax
	// no need to mention length => we can add as many values as we want and it automatically expands the memory
	var fruitList = []string{"Apple", "Mango", "Banana"}
	fmt.Printf("Type of fruitList is %T \n", fruitList)

	// adding new values => append(nameOfSlice, add the values separated by commas) => items are appended at the end
	fruitList = append(fruitList, "Peach", "Pear")
	fmt.Println("Fruit List is", fruitList)

	// for slicing we can also use append mtd as 
	// last range is always non-inclusive
	fruitList = append(fruitList[1:3])
	fmt.Println("Updated List is", fruitList)

	// another synatx - using make() memory mgtm mthd
	// make(type_of_data, size) => the two arguments
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 456
	highScores[3] = 867
	// highScores[4] = 222 => this won't work bcz memory is not allocated
	// but this works bcz append automatically reallocates the memory 
	highScores = append(highScores, 555, 666, 777)

	fmt.Println("High Scores are", highScores)

	// sort is a package so sort thrugh all different types of data structures
	// Ints mthd sort ints in increasing order
	// IntsAreSorted mtd checks if the ints are sorted or not and returns a boolean value
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("Sorted high scores", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// remove a value from slice based on index => very important 
	var courses = []string{"react", "vue", "angular", "javascript", "python"}
	fmt.Println("Courses", courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Updated Courses", courses)
}