package main

import (
	"fmt"
	"sync"
	"time"
)

func processImage(ImageURL string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Processing image: %v\n", ImageURL)
	time.Sleep(3 * time.Second)
	fmt.Printf("Image processing completed: %v\n", ImageURL)
}

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	images := []string{
		"https://example.com/image1.jpg",
		"https://example.com/image2.jpg",
		"https://example.com/image3.jpg",
		"https://example.com/image1.jpg",
	}

	for _, image := range images {
		wg.Add(1)
		go processImage(image, &wg)
	}
	wg.Wait()
	fmt.Printf("Elapsed time: %v\n", time.Since(now))
}