package main

import (
	"fmt"
	"sync"
	"math/rand"
	"runtime"
	"sync/atomic"
)

var wg sync.WaitGroup

func parallelStuff(mutexPtr *int32, mutexPtr2 *int32, sum *int, thrd int){

	//here just to make thread take time
	var intArr []int
	for i:=0; i<10000; i++ {

		intArr=append(intArr, rand.Intn(256))	
	}
	for i :=0; i < 10000; i++ {
		
		locked := atomic.CompareAndSwapInt32(mutexPtr, 0, 1)
		for !locked {
			locked = atomic.CompareAndSwapInt32(mutexPtr, 0, 1)
		}
		locked2 := atomic.CompareAndSwapInt32(mutexPtr2, 0, 1)
		for !locked2 {
			locked2 = atomic.CompareAndSwapInt32(mutexPtr2, 0, 1)
		}
			
		
		//pretend this code required both mutex locks
		*sum=*sum+intArr[i]
		
		*mutexPtr=0
		*mutexPtr2=0
	}
	wg.Done();
}

func main() {
	runtime.GOMAXPROCS(30)
	mutexPtr := new(int32)
	*mutexPtr = 0
	mutexPtr2 := new(int32)
	*mutexPtr2 = 0

	sum1 := new(int)
	sum2 := new(int)
	*sum1 = 0
	*sum2 = 0
	
	wg.Add(2)
	
	go parallelStuff(mutexPtr, mutexPtr2, sum1, 1)
	go parallelStuff(mutexPtr2, mutexPtr, sum2, 2)
	wg.Wait()
	fmt.Println(*sum1)
	fmt.Println(*sum2)
	
}