// handling the get request
// sending data in json format
// sending data in url-encoded format

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
)

func main()  {
    fmt.Println("Web Verb - GET Request")

    // PerformGetRequest()
    // PerformJsonPostRequest()
    PerformPostFormRequest()

}

func PerformGetRequest()  {
    const myurl = "http://localhost:8000/get"

    reponse, err := http.Get(myurl)

    checkNilError(err)

    defer reponse.Body.Close()

    fmt.Println("Status Code:", reponse.StatusCode)
    fmt.Println("Content Length:", reponse.ContentLength)
    
    // reading response => ioutil
    content, _ := ioutil.ReadAll(reponse.Body)

    // instead of this format, we can also use strings pkg which has a Builder mthd for string conversion
    // A Builder is used to efficiently build a string using Write methods. It minimizes memory copying.
    // both the methods work the same
    fmt.Println("Content is:", string(content))
    var responseString strings.Builder
    // with this Write mtd, the content is now in responseString which is in the form of bytes
    byteCount, _ := responseString.Write(content)
    fmt.Println("Byte Count:", byteCount)
    fmt.Println("Content is:", responseString.String())
}

func PerformJsonPostRequest()  {
    const myurl = "http://localhost:8000/post"

    // creating json data - json payload
    // NewReader() allows us to create new reading format
    requestBody := strings.NewReader(`
        {
            "coursename":"Let's go with golang",
            "price": 0,
            "platform":"learnCodeOnline.in"
        }
    `)

    // Post() params => url, content-type, data
    response, err := http.Post(myurl, "application/json", requestBody)

    checkNilError(err)
    defer response.Body.Close()

    content, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(content))
}

func PerformPostFormRequest()  {
    const myurl = "http://localhost:8000/postform"

    // creating form data
    // all the data sent via post request is accessible by url.Values
    data := url.Values{}
    // Add() => adds data
    data.Add("firstname", "noor")
    data.Add("lastname", "poonia")
    data.Add("email", "noor@gmail.com")

    // PostForm() => post form url encoded data
    response, err := http.PostForm(myurl, data)

    checkNilError(err)

    defer response.Body.Close()

    content, _ :=  ioutil.ReadAll(response.Body)

    fmt.Println("Content is:", string(content))

}

func checkNilError(err error)  {
    if err != nil {
        panic(err)
    }
}