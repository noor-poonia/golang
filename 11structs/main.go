package main

import "fmt"

func main()  {

	// structs are really important as they are the alternates of classes 
	// there is no inheritance in golang; no super or parent concepts as well 
	fmt.Println("Structs in Golang")

	// creating a user following the struct type => output is not comma separated 
	john := User{"John", "john@gmail.com", true, 40}
	// output is New user {John john@gmail.com true 40}
	fmt.Println("New user", john)
	// %v => placeholder for printing values => %+v prints in more detail 
	// output is New user details are: {Name:John Email:john@gmail.com Status:true Age:40}
	fmt.Printf("New user details are: %+v\n", john)
	fmt.Printf("New user name is: %v\n", john.Name)
}

// syntax 
// User as capital U => bcz this is like a class or a template which needs to be exported 
// the properties also have 1st letter capital indicating that they can be used from anywhere 
type User struct {
	Name	string
	Email	string
	Status	bool 
	Age	int
}