package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com", "http://stackoverflow.com", "http://www.youtube.com/", "http://github.com/",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	for link := range channel {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, channel)
		}(link)
	}
}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " Is down : ", err)
		c <- link
		return
	}
	c <- link
	fmt.Println("Status code for ", link, " is : ", resp.StatusCode)
}
