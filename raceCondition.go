package main

import (
	"fmt"
	"sync"
)

/*
Race condition explanation:
When you run the go file, you will run two go routines concurrently.
The go routines share a common variable called a counter. One of the go routines decrement the counter,
whereas the other one increments it. If there is a context switch right when the increment routine is running,
the decrement routine would decrement the counter, but when the increment routine gets back the control,
it would be unaware of the decrement and go ahead with incrementing and overwrite the value of the counter.
So, if the initial value is 0, the final result would be 1, instead of remaining 0.
*/

const NumAction int = 10000

var (
	counter int
	wg      sync.WaitGroup
)

func increment() {
	for i := 0; i < NumAction; i++ {
		counter++
	}

	wg.Done()
}

func decrement() {
	for i := 0; i < NumAction; i++ {
		counter--
	}
	wg.Done()
}

func main() {

	wg.Add(2)

	go increment()
	go decrement()

	wg.Wait()
	fmt.Println("Counter:", counter)
}
