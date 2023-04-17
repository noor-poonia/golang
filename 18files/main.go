package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golang")

	content := "This needs to go in a file"

	// Create(filePath) is a func used to create file 
	// comma ok syntax used bcz in some cases file maybe created or their might be some issues as well 
	file, err := os.Create("./file.txt")

	checkNilError(err)

	// for writing in the file 
	// io is another package -> writing in the files 
	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length is", length)

	// recommended to write defer keyword
	defer file.Close()

	readFile("./file.txt")
}

func readFile(fileName string) {
	// ioutil -> for reading a file
	// data read from the file is in data byte format and not string format
	dataByte, err := ioutil.ReadFile(fileName)

	checkNilError(err)

	// conversion of dataByte into string
	fmt.Println("Text data inside the file is: \n", string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		// panic() shuts down the program and throws the error 
		panic(err)
	}
}
