package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
		"https://dfsfdsfsdfesrew.com",
	}

	// Create a new channel
	c := make(chan string)

	for _, link := range links {
		// Use Go Routine
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c) // (<-c) waits for value to be sent. This is blocking the thread.
	}

}

// channel of string
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep it's up" // c <- "xxx" sends message to channel
}
