package main

import (
	"fmt"
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
		"https://dfsfdsfsdfesrew.com",
	}

	// Create a new channel
	c := make(chan string)

	for _, link := range links {
		// Use Go Routine
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c) // (<-c) waits for value to be sent. This is blocking the thread.
	// }

	// // Infinite loop after finish checking one link
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Wait for channel to return value, assign to l and run functions
	// Same as infinite loop. But more understandable
	for l := range c {
		// time.Sleep(5 * time.Second) // This will block main routine. NG

		go func(link string) { // Function Literal. Used to pass function as argument
			time.Sleep(5 * time.Second)
			checkLink(link, c) // Pass l as value to the Function Literal to copy it (by value).
		}(l)
		// go checkLink(l, c)
	}
}

// channel of string
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link // c <- "xxx" sends message to channel
}
