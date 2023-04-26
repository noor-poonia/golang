package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev/"

func main()  {
	fmt.Println("web Requests in golang")

	// https => net/http package is used for handling web requests 
	// Get mtd takes one argument => url
	response, err := http.Get(url)

	checkNilError(err)

	// the type is *http.Response
	fmt.Printf("Type of response is: %T\n", response)

	// caller's responsibility to close the connection
	defer response.Body.Close()

	// reading the response => majority by ioutil
	dataByte, err := ioutil.ReadAll(response.Body)

	checkNilError(err)

	fmt.Println("Data is:", string(dataByte))

}

func checkNilError(err error)  {
	if err != nil {
		panic(err)
	}
}