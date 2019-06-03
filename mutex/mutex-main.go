package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	intPointer := []int{0}
	
	wg.Add(1000)
	
	for i :=0; i < 1000; i++ {
		go func(iP [] int) {
			
			iP[0]=iP[0]+1;
			
			wg.Done();
		}(intPointer)
	}
	
	wg.Wait()
	fmt.Println(intPointer[0])
	
}