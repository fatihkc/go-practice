package main

import (
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	time.Sleep(time.Second * 5)
	_, err := http.Get(link)
	if err != nil {
		c <- link + " might be down!"
		return
	} else {
		c <- link + " is up!"
		return
	}
}
