package main


import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)
func main() {
	start:=time.Now()
	var intArr []int
	sum:=0
	for i:=0;i<32768; i++ {
		intArr=append(intArr,rand.Intn(256))
	}
	sort.Ints(intArr);
	for i:=0; i<100000; i++ {
		for j:=0; j<30000; j++ {
			if intArr[j]>128 {
				sum=sum+intArr[j]
			}
		}
	}
	t:=time.Now()
	fmt.Println(t.Sub(start));
	fmt.Println(sum);
}