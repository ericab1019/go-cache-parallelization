package main

import (
	"fmt"
	"time"
	"math/rand"
)

//if run with j first, takes 5 seconds on average on benchmark computer. Run with i first, takes 1 second. 
func main() {
	start:=time.Now()
	
	const loopsize = 10000
	
	var intArr [loopsize][loopsize]int
	
	var readArr []int
	
	for i:=0;i<loopsize; i++ {
		readArr=append(readArr, rand.Intn(50))
	}
		
	for j:=0; j<loopsize; j++ {
		
		for i:=0; i<loopsize; i++ {
			intArr[i][j] += i+readArr[j]
		}
	}
	t:=time.Now()
	fmt.Println(t.Sub(start));
}

