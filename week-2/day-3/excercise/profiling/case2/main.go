package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"runtime/pprof"
	"time"
)

func doSomething() {
	for i := 0; i < 10000; i++ {
		rand.Intn(10000)
	}
}

func main() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	go func() {
		fmt.Println(http.ListenAndServe("localhost:8080", nil))
	}()
	for i := 0; i < 1000; i++ {
		go doSomething()
	}
	time.Sleep(5 * time.Second)
}