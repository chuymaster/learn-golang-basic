// Assignment: https://www.udemy.com/go-the-complete-developers-guide/learn/v4/t/practice/9178/introduction

package main

import (
	"strconv"
)

func main() {

	// create a slice of ints from 0 through 10
	intsSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Iterate through the slice with a for loop.  For each element, check to see if the number is even or odd.
	for _, value := range intsSlice {

		// If the value is even, print out "even".  If it is odd, print out "odd"

		sValue := strconv.Itoa(value)

		if value%2 == 0 {
			println(sValue + " is even")
		} else {
			println(sValue + " is odd")
		}
	}

}
