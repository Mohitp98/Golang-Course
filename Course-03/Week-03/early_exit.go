package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//only main routine is printed (without sleep)

	// goroutine
	go fmt.Println("This is from goroutine call")

	//main routine
	fmt.Println("This is normal call")
	time.Sleep(1000 * time.Microsecond)

	/*
		// with waitGroup
		wg.Add(1)			// increase counter ++
		go simpleFunct()	// execute thread
		wg.Wait()			// wait for execution

		//main routine
		fmt.Println("This is normal call")
	*/
}

func simpleFunct() {
	go fmt.Println("This is from goroutine call") // call goroutine
	wg.Done()                                     // decrement counter --
}
