package main

import (
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "strconv"
    "time"

    "github.com/gorilla/mux"
)

// in some cases we can also have multiple models interacting with each other
// creating a model for course - usually in a different file
type Course struct {
    CourseId string `json:"courseid"`
    CourseName string `json:"coursename"`
    CoursePrice int `json:"price"`
    Author *Author `json:"author"`
    // type is now pointer
    // Author Author => this will work but its like passing a copy of it and we can't also reference as of now bcz it hasn't been initialized
}

type Author struct {
    Fullname string `json:"fullname"`
    Website string `json:"website"`
}

// creating db using slice for now
var courses []Course

// middlewares or helpers - usually in a different file 
func (c *Course) IsEmpty() bool {
    // return c.CourseId == "" && c.CourseName == ""
    return c.CourseName == ""
}

func main()  {
    fmt.Println("API in Golang")

    // mux is used to deal with routers and dealing with controllers
    r := mux.NewRouter()

    // seeding -> process of populating database with initial set of data
    courses = append(courses, Course{
        CourseId: "2", 
        CourseName: "Reactjs",
        CoursePrice: 299,
        // since author is pointer type, we need to pass it as a reference
        Author: &Author{
            Fullname: "John",
            Website: "lco.dev",
        },
    })
    courses = append(courses, Course{
        CourseId: "4", 
        CourseName: "MERN Stack",
        CoursePrice: 199,
        // since author is pointer type, we need to pass it as a reference
        Author: &Author{
            Fullname: "Jack",
            Website: "go.dev",
        },
    })

    // routing 
    r.HandleFunc("/", serveHome).Methods("GET")
    r.HandleFunc("/courses", getAllCourses).Methods("GET")
    r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
    r.HandleFunc("/course", createOneCourse).Methods("POST")
    r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
    r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")


    // listen to port 
    log.Fatal(http.ListenAndServe(":4000", r))
}

// controller basically handles the situation where route is involved => it defines what to do when a certain route is hit
// controllers -> usually in a separate file 
// serve home route => 2args => writer w and reader r => this func defies what happens when we are in the home route
// reader => u get the value from sender
// writer => write the response for that
// these arguments are mandatory and should be passed in the same order
func serveHome(w http.ResponseWriter, r *http.Request)  {
    w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("Get all courses")
    // setting the header -> key value pair
    w.Header().Set("Content-Type", "application/json")
    // arg inside Encode() will be treated as json value
    json.NewEncoder(w).Encode(courses)
}

// in this case, we will use the reader to get the course id - its the unique value to idetify the particuar course
func getOneCourse(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("Get one course")
    w.Header().Set("Content-Type", "application/json")
    
    // grab id from request
    // params will hold all the data coming in 
    // mux has a function to get all the variablles/params coming in -> directly from the url
    // Vars() takes the reader as argument and returns all the params/variables coming in 
    params := mux.Vars(r)
    // params is key - value pair

    fmt.Println("Params is:", params)
    fmt.Printf("Type of params is %T\n", params)

    // loop through the courses, find the matching id and return the response
    for _, course := range courses {
        if course.CourseId == params["id"] {
            json.NewEncoder(w).Encode(course)
            return 
        }
    }
    json.NewEncoder(w).Encode("No course found with the given id")
    return
}

func createOneCourse(w http.ResponseWriter, r *http.Request)  {
    // this time we are receiving json, so we will be decoding it instead of encoding

    fmt.Println("Create one course")
    w.Header().Set("Content-Type", "application/json")

    // there are different cases we need to take care of 
    // 1.) Body is empty
    if r.Body == nil {
        json.NewEncoder(w).Encode("Please send some data")
    }

    // 2.) Body = {}
    var course Course
    // decode the value bcz data is expected to be in a particular structure or format
    _ = json.NewDecoder(r.Body).Decode(&course)
    // this works smthg like r.Body is the value and what it should be destructured on
    if course.IsEmpty() {
        json.NewEncoder(w).Encode("No data inside JSON")
        return 
    }

    // TODO: check only if title is duplicate -> loop, title matches with course name, Json response course already exists 
    for _, c := range courses {
        if course.CourseName == c.CourseName {
            json.NewEncoder(w).Encode("Course Already Exists")
        }
    }

    // generate unique id, string conversion, append course into courses
    rand.Seed(time.Now().UnixNano())
    // converting this random integer value into string
    course.CourseId = strconv.Itoa(rand.Intn(100))
    courses = append(courses, course)
    json.NewEncoder(w).Encode(course)
    return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request)  {
    // identify the specific course to update by the unique course id
    // we need all the values that need to updates

    fmt.Println("Update one course")
    w.Header().Set("Content-Type", "application/json")

    // take id from request
    params := mux.Vars(r)
    // this time we need 2 values -> id from url and json data from body

    // loop through, get id, remove that data, add it again with my ID (id from params)

    for index, course := range courses {
        if course.CourseId == params["id"] {
            courses = append(courses[:index], courses[index+1:]...)
            var course Course
            _ = json.NewDecoder(r.Body).Decode(&course)
            // over writting the values
            course.CourseId = params["id"] 
            courses = append(courses, course)
            json.NewEncoder(w).Encode(course)
            return
        }
    } 
    // TODO: send a response when id is not found
    json.NewEncoder(w).Encode("No ID found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("Delete one course")
    w.Header().Set("Content-Type", "application/json")

    // id -> to know which course to delete
    params := mux.Vars(r)
    
    // loop, id match, remove (index, index+1)
    for index, course := range courses {
        if course.CourseId == params["id"] {
            courses = append(courses[:index], courses[index+1:]...)
            json.NewEncoder(w).Encode("Course successfully deleted")
            return
        }
    }
    // TODO: send a response if id is not found
    json.NewEncoder(w).Encode("No ID found")
}

// usually in golang for api calling, the function names are not like getOneCourse or getAllCourses instead they are just courses and based the 
// request type (get, post) its handled