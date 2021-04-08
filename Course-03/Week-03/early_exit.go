package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//var wg sync.WaitGroup

	// goroutine
	go fmt.Println("This is from goroutine call")

	//main routine
	fmt.Println("This is normal call")
	//only main routine is printed (without sleep)
	time.Sleep(1000 * time.Microsecond)

	/*
		// with waitGroup
		wg.Add(1)			// increase counter ++
		go simpleFunct(&wg)	// execute thread
		wg.Wait()			// wait for execution

		//main routine
		fmt.Println("This is normal call")
	*/
}

func simpleFunct(wg *sync.WaitGroup) {
	go fmt.Println("This is from goroutine call") // call goroutine
	defer wg.Done()                               // decrement counter --
}
