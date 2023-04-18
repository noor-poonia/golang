package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main()  {
    fmt.Println("Mod in Golang")

    // go.mod files initializes the program and includes the go version on which our program is based on 
    // versioning is done as :- 1.0.0 -> new release, 2.0.0 -> major release, 1.1.0 -> minor release, 1.0.1 -> patch release => semantic versioning
    // gorilla/mux => routing feature in golang
    // go.sum => compares the hash of original repo with the repo we are using
    // go mod verify => goes into go.mod to check the number of pkgs and then into go.sum to verify the hashes
    // go mod tidy => tidies up all the dependant libraries -> expensive operation
    // go mod graph => list the modules and the dependancies
    // go mod edit => t edit the go.mod file
    // go mod vendor -> creates a vendor folder -> smthg like node_modules
    // go run -mod=vendor main.go => it will first look into vendor folder and then web for depandencies 

    greeter()

    r := mux.NewRouter()
    r.HandleFunc("/", serveHome).Methods("GET")

    log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
    fmt.Println("Mod Users")
}

func serveHome(w http.ResponseWriter, r *http.Request)  {
    w.Write([]byte("<h1>Welcome to Golang</h1>"))
}