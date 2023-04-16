package main

import "fmt"

func main()  {
	fmt.Println("Loops in Golang")

	// slice
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// looping through slices
	// its similar to for (int i = 0; i < len(days); i++) => ++d doesn't exist
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// range automatically loops through the slice
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for each loop => used a lot 
	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	// almost similar to while loop
	// and goto stmt also 
	rogueValue := 1
	for rogueValue < 7 {
		if rogueValue == 2 {
			goto lco
		}
		if rogueValue == 5 {
			// break
			rogueValue++
			continue
		}
		fmt.Println("value is:", rogueValue)
		rogueValue++
	}

	lco:
		fmt.Println("Jumping at LearnCodeOnline.in")
}