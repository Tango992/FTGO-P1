package main

import (
	"fmt"
	"math/rand"
)


func main() {
	i := rand.Intn(65) + 1
	ch := make(chan int, 1)
	defer close(ch)
	
	ch <- i
	iteration, result := factorial(ch)
	fmt.Printf("%v factorial: %v\n", iteration, <-result)
}


func factorial(ch <-chan int) (int, chan uint64){
	chOut := make(chan uint64, 1)
	defer close(chOut)
	n := <-ch
	
	var result uint64 = 1
	for i := 1; i <= n; i++ {
		result *= uint64(i)
	}
	chOut <- result
	return n, chOut
}
