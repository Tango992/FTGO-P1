package main

import (
	"fmt"
	"net/http"
)

func main() {
	// switchCase()
	statistic()
}

func statistic() {
	var sum int
	slice := []int{1,2,3,4,3,2,1,2,3,4,3,2,1}
	length := len(slice)

	for _, i := range slice {
		sum += i
	}

	avg := float32(sum) / float32(length)
	fmt.Printf("%.3f\n", avg)
}

func switchCase() {
	response, err := http.Get("http://www.linkedin.com")

	if err != nil {
		fmt.Println(err)
		return
	}

	switch response.StatusCode {
	case 200:
		fmt.Println("OK")
	case 404:
		fmt.Println("Not found")
	case 500:
		fmt.Println("Internal Server Error")
	case 502:
		fmt.Println("Bad Gateway")
	default:
		fmt.Printf("Received HTTP Status code: %v\n", response.StatusCode)
	}
}