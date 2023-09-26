package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	Url := "https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4"
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			iteration := fmt.Sprint(i)
			filename := fmt.Sprintf("video%v.mp4", iteration)
			err := DownloadFile(filename, Url)
			if err != nil {
				fmt.Println("Error downloading file:", err)
				return
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("All files have been downloaded")
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	fmt.Println("Downloaded:", filepath)
	return err
}