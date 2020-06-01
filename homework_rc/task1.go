package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)
	var mutex = &sync.Mutex{}

	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}
