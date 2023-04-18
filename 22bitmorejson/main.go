package main

import (
    "encoding/json"
    "fmt"
)

// struct name starting with lowercase => not making it public
type course struct {
    // this makes the end result change to coursename instead of Name
    Name     string `json:"coursename"`
    Price    int 
    Platform string `json:"website"`
    // - means this field shouldn't be reflected to who ever is using this api
    Password string `json:"-"`
    // omitmpty => if the value is null/nil don't show it at all
    // don't put space after ,
    Tags     []string `json:"tags,omitempty"`
}

func main() {
    fmt.Println("JSON in Golang")

    // EncodeJson()
    DecodeJson()
}

// encoding of json => converting any type of data (string, ket-value pair, etc) into valid json
func EncodeJson()  {
    // the type here is the struct courses
    lcoCourses := []course{
        {"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
        {"MERN Bootcamp", 199, "LearnCodeOnline.in", "def123", []string{"full-stack", "js"}},
        {"Angular Bootcamp", 299, "LearnCodeOnline.in", "ghi123", nil},
    }

    // goal => package this data as json data
    // json => package, Marshal() => kind of implementatiopn of json => returns json encoding => argument is interface (always) => kind of an alternate word for struct
    // finalJson, err := json.Marshal(lcoCourses)
    // this mtd makes it easier to read
    // arguments => interface, prefix(word at the start of every new line), indent(here \t => tab)
    finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

    checkNilError(err)

    // this prints in byte format
    // fmt.Println("JSON data:", finalJson)
    // this prints an array => hard to read, able to see password, last Tags shows null instead of nil
    fmt.Printf("%s\n", finalJson)
}

// decoding json => consuming json data 
func DecodeJson()  {
    jsonDataFromWeb := []byte(`
        {
            "coursename": "ReactJS Bootcamp",
            "Price": 299,
            "website": "LearnCodeOnline.in",
            "tags": ["web-dev", "js"]
        }
    `)

    var lcoCourse course

    // to check if json data is valid or not 
    // returns a boolean value
    checkValid := json.Valid(jsonDataFromWeb)

    if checkValid {
        fmt.Println("JSON was valid")
        // passing a reference instead of passing a copy
        json.Unmarshal(jsonDataFromWeb, &lcoCourse)
        // %#v -> to print interfaces
        fmt.Printf("%#v\n", lcoCourse)
    } else {
        fmt.Println("JSON was not valid")
    }

    // there are some cases where you want to add data to key value pair 
    // value type is interface bcz we don't know the type of data coming from web 
    var myOnlineData map[string]interface{}
    json.Unmarshal(jsonDataFromWeb, &myOnlineData)
    fmt.Printf("%#v\n", myOnlineData)

    for key, value := range myOnlineData {
        fmt.Printf("key is %v and value is %v and type is %T\n", key, value, value)
    }
}

func checkNilError(err error) {
    if (err != nil) {
        panic(err)
    }
}