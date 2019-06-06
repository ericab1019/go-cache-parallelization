package main

import (
	"fmt"
	"time"
	"sync"
//	"sync/atomic"
)

type shared struct {
    x uint32
    y uint32
}
var share=shared{2,3}

var wg sync.WaitGroup
func sumA() {
	for i:=0; i<10000000; i++ {
		share.x=share.x+1
		//atomic.AddUint32(&share.x, 1)
	}
	wg.Done()
} 	

func sumB(){
	for i:=0; i<10000000; i++ {
		share.y=share.y+1
		//atomic.AddUint32(&share.y, 1)
	}
	wg.Done();
}
	
	
//without the padding, processor A and processor B will keep invalidating each other's cache line forcing cache misses. 
// Code shows a difference in performance between padding and non-padding without the atomic, but atomic makes difference more noticeable
func main() {
	start:=time.Now()
	wg.Add(2)
	go sumA()
	go sumB()
	//go func() {sumB()}()
	wg.Wait()
	t:=time.Now()
	fmt.Println(t.Sub(start));
}