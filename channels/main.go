package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.amazon.com/",
		"https://stackoverflow.com/",
		"https://golang.org/",
	}

	c := make(chan string)

	for _, link := range links {

		// go only works in front of function calls
		go checkLink(link, c)
	}

	for l := range c {
		// <-c is a blocking call that waits for the message to be passed from checkLink
		//fmt.Println(<-c)
		go func(link string) {
			time.Sleep(3 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Request to", link, "failed")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	//c <- "It's up"
	c <- link
}
