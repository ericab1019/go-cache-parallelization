package main

import (
	"fmt"
	"time"
	"math/rand"
)

//with blocking by blocksize runs in about 10s, without blocksize lose spatial locality and takes 17s
func main() {
	//blocksize:=8
	start:=time.Now()
	
	loopsize := 40000
	
	var intArr []int
	
	var readArr []int
	
	for i:=0;i<loopsize; i++ {
		intArr=append(intArr, 0)
	}
	for i:=0;i<loopsize; i++ {
		readArr=append(readArr, rand.Intn(50))
	}
	
	//for j2:=0; j2<loopsize; j2=j2+blocksize {
		
	for i:=0; i<loopsize; i++ {
		
		for j:=0; j<loopsize; j++ {
			intArr[i] += i+readArr[j]
		}
	}
		//for j:=j2; j<j2+blocksize; j++ {
			//intArr[i] += i+readArr[j]	
		//}
	//}
	t:=time.Now()
	fmt.Println(t.Sub(start));
}

