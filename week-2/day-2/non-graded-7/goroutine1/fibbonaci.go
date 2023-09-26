package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// with goroutine: 500-600 micro seconds
	// without goroutine: 190-250 micro seconds
	now := time.Now()
	var wg sync.WaitGroup
	for i := 1; i <= 200; i++ {
		wg.Add(1)
		go func(i int){
			fib(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Since(now))
}

func fib(n int){
	var a uint64 = 0
	var b uint64 = 1
	var c uint64
	if n == 0 {
		fmt.Println(a)
		return
	} else if n == 1{
		fmt.Println(b)
		return
	} else {
		for i := 2; i <= n; i++ {
			c = a + b
			a = b
			b = c
		}
		fmt.Println(b)
		return
	}
}