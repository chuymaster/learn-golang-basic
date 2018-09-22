// Interface assignment 2 https://www.udemy.com/go-the-complete-developers-guide/learn/v4/t/practice/9958/introduction

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// fmt.Println(os.Args)
	argument := os.Args[1]
	fmt.Println(argument)

	file, err := os.Open(argument) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
