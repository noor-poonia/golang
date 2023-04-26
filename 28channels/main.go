package main

import (
	"fmt"
	"sync"
)

func main() {
	// channels are the way in which the multiplle go routines can talk to each other
	// they are like a pipeline through which multiple go routines interact
	// it allows us to pass a value if somebody is listening to it 
	// you still won't know what is happening or how much tme will it take but there might be some sort of a signal 
	fmt.Println("Channels in Golang")

	// creating a channel -> int is basically the type in which the goroutines interact among each other 
	myChannel := make(chan int, 2)
	// this {int} is a buffer value => consuming one value => if there are more values it will only consider the 1st one and won't show any error 

	// this actually won't work without a go routine so for that we need to 1st create a wait group
	wg := &sync.WaitGroup{}

	// entering a value in channel 
	// myChannel <- 5
	// printing values => taking the value out thats why the arrow placing is like that 
	// fmt.Println("channel:", <-myChannel)
	// also with channels 1st we need to receive the value the put values into it 

	wg.Add(2)
	// responsible for receiving a value
	go func(ch <-chan int, wg *sync.WaitGroup)  {
		// to solve the proble with 0, we can do 
		val, ischannelOpen := <-myChannel
		fmt.Println(val)
		fmt.Println(ischannelOpen)
		// fmt.Println(<-myChannel)
		// fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel, wg)

	// responible fofr putting values into channels - send only
	go func(ch chan<- int, wg *sync.WaitGroup)  {
		myChannel <- 5 // for this we have a listener 
		myChannel <- 6 // for this the existing listener won't work, we need to add the listener again 
		// but if we don't want more listeners, we can use buffered channel 
		// listening to a closed channel returns a 0 but it may be problematic if that 0 is also the value passed 
		close(myChannel)
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}