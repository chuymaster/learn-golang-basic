package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson"} // This shows how to initialize the struct
	println(alex.firstName)
	fmt.Println(alex)
	fmt.Printf("%+v", alex) // Print format. Print all fields and values.

	var john person // This automatically assign Zero value to all variables in struct
	fmt.Println(john)

	john.firstName = "John"
	john.lastName = "Doe"
	fmt.Println(john)
}
