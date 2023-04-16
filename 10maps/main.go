package main 

import "fmt"

func main()  {

	// maps => key - value pair => key and value can be anything 
	// retrieval are very fast 
	fmt.Println("Maps in Golang")

	// syntax 
	// we are using make() bcz new() gives a lot of errors 
	// argument is the type "map[key-type]value-type" 
	languages := make(map[string]string)

	// adding the values 
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	// output is Languages are map[JS:Javascript PY:Python RB:Ruby] => not comma separated 
	fmt.Println("Languages are:", languages)

	// retrieving one value => mapName[key]
	fmt.Println("JS shorts for:", languages["JS"])

	// deleting values => arguments => mapName, key => can be used for slices also 
	delete(languages, "RB")
	fmt.Println("Languages are:", languages)

	// loops are interesting in golang => key, value := range is again the comma ok syntax 
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
	
}