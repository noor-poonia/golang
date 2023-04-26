package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/noor-poonia/mongoApi/router"
)

func main() {
	// refer go.mongodb.org
	// in golang -> root directory can have only one go file
	fmt.Println("MongoDB in Golang")

	r := router.Router()
	fmt.Println("Server is getting started....")
	fmt.Println("Listening at port 4000...")
	log.Fatal(http.ListenAndServe(":4000", r))
}
