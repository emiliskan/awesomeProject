package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	urls := []string{
		"http://google.com",
		"http://facebok.com",
		"http://golang.org",
		"http://vk.com",
		"http://ya.ru",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkLink(url, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(url string, c chan string) {
	resp, _ := http.Get(url)
	if resp.StatusCode != 200 {
		c <- url
		fmt.Println(url, "do not work!")
		return
	}

	fmt.Println(url, "work!")
	c <- url
}
