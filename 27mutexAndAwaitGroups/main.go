package main

import (
	"fmt"
	"sync"
)

// on the surface level, it looks like its working fine 
// ends with exit status 66

func main()  {
	fmt.Println("Race condition in Golang")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	mut.RLock()
	var score = []int{0}
	mut.RUnlock()

	wg.Add(3) // defining that there are a total of 3 goroutines

	// anonymous function -> goroutines
	go func(wg *sync.WaitGroup,  m *sync.RWMutex) {
		fmt.Println("1st goroutine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("2nd goroutine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("3nd goroutine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("3nd goroutine")
		// mut.Lock()
		// score = append(score, 3)
		// mut.Unlock()
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}