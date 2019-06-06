package main

import (
	"fmt"
	"sync"
	"time"
)

//This function exemplifies why parallelism (or in this example moreso concurrency) speeds up blocking calls like sleep. 
//Would also speed up IO
func main() {

	var  wg sync.WaitGroup
	wg.Add(1000)
	for i:=0; i<1000; i++ {
		go func() {
			time.Sleep(1000*time.Millisecond)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("done")
}