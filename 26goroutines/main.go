// concurrency vs parallelism -> imprtant concept
// concurrency -> allowing yourself to do multiple tasks but not at the same time - doing them one at a time while switching from one task to another
// parrallelism -> multitasking (in simple words) -> basically not switching from one task to another

// goroutines -> way to achieve parallelsim -> often compared to threads
// goroutines > managed by Go runtime and is flexible stack - 2kb
// thread -> managed by os and is fixed stack - 1mb
// slogan for golang -> do not communicate by sharing memory; instead share memory by communicating

package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

var signals = []string{"test"}

// wait grp -> modified version of time.sleep
// 3 jobs -> add, done, wait
// 1. add -> as soon as go routine is created add that in the wait grp which is responsible for mgmt
// 2. done -> whenever the job is done, it should automatically call done 
// 3. wait -> if the job is not done, it should cal wait 
// NOTE: usually these are pointer
var wg sync.WaitGroup

// mutex -> mutual exclusion lock -> if one go routine is using some memory then it won't allow any other process to use it -> RWMutex is also similar
// output wise not much difference but internally its important
// usually a pointer 
var mut sync.Mutex

func main()  {
	fmt.Println("Goroutines in Golang")

	// goroutine is simply created by adding the keyword go
	// it fires up (launch) the thread and later the time pacage ensures when to bring back the thread for execution 
	// instead of using time package, there are packages available for this particular thing -> sync package
	// go greeter("hello")
	// greeter("world")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		// this adds the number of how many goroutines are launched
		wg.Add(1)
	}

	// this is usually at the end of the function -> doesn't allow the function to finish 
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string)  {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// this process is a bit a slow then expected -> it can be done faster by introducing some threads
func getStatusCode(endpoint string) {
	// passing a signal that the job is done
	defer wg.Done()

	result, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
	}

	// placeholder %d -> for integers, %s -> for string 
	
}