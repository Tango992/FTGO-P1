package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		fmt.Println("Mulai mengirim pesan...")
		ch <- "Hello dari unbuffered Channel!"
		fmt.Println("Pesan terkirim...")
	}()

	time.Sleep(2 * time.Second)
	
	go func() {
		message := <- ch
		fmt.Println("Pesan diterima:", message)
	}()
	time.Sleep(1 * time.Second)
}