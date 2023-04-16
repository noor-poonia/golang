package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("Switch and case in Golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is", diceNumber)

	// syntax
	// there is no autmatic fall through happeing or simply there's no break stmt
	// so we need to explicitly mention "fallthrough" after the working
	// fallthrough basically runs that case and the case after that as well
	switch diceNumber {
	case 1: 
		fmt.Println("Dice value is 1 and you can open")
	case 2: 
		fmt.Println("You can move 2 spots")
	case 3: 
		fmt.Println("You can move 3 spots")
		fallthrough
	case 4: 
		fmt.Println("You can move 4 spots")
	case 5: 
		fmt.Println("You can move 5 spots")
		// fallthrough
	case 6: 
		fmt.Println("You can move 6 spots and roll dice again")
	default: 
		fmt.Println("What was that!!!")
	}
}