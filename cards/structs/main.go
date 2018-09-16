package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
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

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		}, // Multiline declaration requires comma at the end
	}

	jim.print()
	jim.updateName("Jimmy") // This will not work because GO is pass by value language
	jim.print()

	// *** Everything in GO is passed by value. Everything is copied before passing. **

	jimPointer := &jim // & creates a pointer to jim
	jimPointer.updateNameUsingPointer("Jimmy")
	jim.print()

	jane := person{
		firstName: "Jane",
		lastName:  "Pork",
		contactInfo: contactInfo{
			email:   "jane@gmail.com",
			zipCode: 94000,
		}, // Multiline declaration requires comma at the end
	}
	jane.print()
	jane.updateNameUsingPointer("Janny") // Shortcut. No need to assign this as pointer because GO converts it automatically.
	jane.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// Start operator change the variable to pointer type
func (p *person) updateNameUsingPointer(newFirstName string) {
	(*p).firstName = newFirstName // (*) turns the pointer to the value inside the memory
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println("")
}
