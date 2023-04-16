package main

import "fmt"

func main()  {

	// not very different from functions 
	// when functions are used in structs, we tend to call them as methods
	fmt.Println("Methods in Golang")

	john := User{"John", "john@gmail.com", true, 40}
	fmt.Println("New user", john)
	john.GetStatus()
	john.NewMail()
	fmt.Println("New user", john)
}

type User struct {
	Name	string
	Email	string
	Status	bool 
	Age	int
}

// syntax is func (struct type) mtdName() {}
func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

// it doesn't actually change the email in the original struct bcz whenever we pass u User like that
// it actually passes a copy & that is why we need pointers => u *User
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}