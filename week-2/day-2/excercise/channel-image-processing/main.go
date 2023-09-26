package main

import (
	"fmt"
	"time"
)

func processImage(i int, ImageURL string, channel chan string) {
	fmt.Printf("Processing image: %v\n", ImageURL)
	time.Sleep(3 * time.Second)
	fmt.Printf("Image processing completed: %v\n", ImageURL)
	channel <- fmt.Sprintf("Iteration = %v", i+1)
}

func main() {
	now := time.Now()
	channel := make(chan string)
	images := []string{
		"https://example.com/image1.jpg",
		"https://example.com/image2.jpg",
		"https://example.com/image3.jpg",
		"https://example.com/image4.jpg",
	}

	for i, image := range images {
		go processImage(i, image, channel)
	}

	for i := 0; i < len(images); i++ {
		fmt.Println(<- channel)
	}
	fmt.Printf("Elapsed time: %v\n", time.Since(now))
}