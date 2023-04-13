package main

import (
	"fmt"
	"time"
)

func main()  {

	// time package => for measuring and displaying time
	// func (Time) Format => commonly used => very crucial and we need to properly specify the values
	fmt.Println("Welcome to Time Study of Golang")

	presentTime := time.Now()

	fmt.Println("Present Time : ", presentTime)
	// this output Present Time : 2023-04-13 11:36:19.952304047 +0530 IST m=+0.000036701

	// .Format() is used for formating the time to make it more understandable
	// the arguments are as follows :- 
	// according to golang documentation, we will always use "01-02-2006" time as standard time for formatting
	// this is MM/DD/YY format
	// to see the current ticking time, we need to pass "15:04:05" as well
	// to see the day also, we need to pass "Monday" as well 
	// of only one argument is there then only that value will be displayed
	// the order of arguments inot important
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// output is 04-13-2023 11:51:58 Thursday

	createdDate := time.Date(2020, time.October, 31, 23, 23, 23, 0, time.UTC)

	fmt.Println("Created Date : ", createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04:05"))
}