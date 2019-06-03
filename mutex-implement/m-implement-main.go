package main

import (
	"fmt"
	"sync"
//	"sync/atomic"
)
// naive implementation locks, or with log statement added will not return 1000, need to use atomic compare and swap commented out lines
func main() {
	
	mutexPtr := new(int)
	*mutexPtr = 0
	
	var wg sync.WaitGroup
	intPointer := []int{0}
	
	wg.Add(1000)
	
	for i :=0; i < 1000; i++ {
		go func(iP [] int, mutexPtr *int) {
		
			locked:=(*mutexPtr==1)
			for locked {
				locked=(*mutexPtr==1)
				fmt.Println(1)
			}
			*mutexPtr=1;
			//locked := atomic.CompareAndSwapInt32(mutexPtr, 0, 1)
			//for !locked {
				//locked = atomic.CompareAndSwapInt32(mutexPtr, 0, 1)
			//}
			
			iP[0]=iP[0]+1;
			*mutexPtr=0;	
			
			wg.Done();
		}(intPointer, mutexPtr)
	}
	
	wg.Wait()
	fmt.Println(intPointer[0])
	
}