package main

import (
	"fmt"
)

func main() {
	chInput := make(chan [2]float64,2)
	chAddOutput := make(chan float64)
	chSubOutput := make(chan float64)
	
	chInput <- [2]float64{5, 3}
	chInput <- [2]float64{5, 3}

	go func() {
		data := <-chInput
		chAddOutput <- data[0] + data[1]
	}()
	go func() {
		data := <-chInput
		chSubOutput <- data[0] - data[1]
	}()

	fmt.Println(<- chAddOutput)
	fmt.Println(<- chSubOutput)
}