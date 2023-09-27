package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func doSomething() {
	for i := 0; i < 10000; i++ {
		rand.Intn(10000)
	}
}

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:8080", nil))
	}()
	for i := 0; i < 1000; i++ {
		go doSomething()
	}
	time.Sleep(200 * time.Second)
}