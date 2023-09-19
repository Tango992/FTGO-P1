package main

import (
	"fmt"
	"net/http"
)

func main() {
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