package main

import (
    "fmt"
    "net/url"
)

// 1. creating fictitious url - try to extract info from it
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
    fmt.Println("Handling URLs in Golang")

    fmt.Println("Url is", myurl)

    // parsing
    result, _ := url.Parse(myurl)

    // Scheme is the words before // => https
    fmt.Println("Result Scheme is", result.Scheme)
    // Host => lco.dev:3000
    fmt.Println("Result Host is", result.Host)
    // path => /learn
    fmt.Println("Result Path is", result.Path)
    // raw query => coursename=reactjs&paymentid=ghbj456ghb
    fmt.Println("Result RawQuery is", result.RawQuery)
    // port => 3000
    fmt.Println("Result Port is", result.Port())

    // Query() stores all the queries in a better format => url.values => basically key, value pair => map[coursename:[reactjs] paymentid:[ghbj456ghb]]
    qparams := result.Query()
    fmt.Printf("Query type is: %T\n", qparams)
    fmt.Println("Query:", qparams)
    fmt.Println("Query course name:", qparams["coursename"])

    for _, value := range qparams {
        fmt.Println("Param is:", value)
    }

    // 2. create info and then try to create a url from it => super easy
    // always pass reference and not copy
    // all the data names is accepted as it is 
    partsOfUrl := &url.URL{
        Scheme: "https",
        Host: "lco.dev",
        Path: "/tutcss",
        RawQuery: "user=john",
    }

    // another conversion format
    anotherURL := partsOfUrl.String()
    fmt.Println("Another URL", anotherURL)

}