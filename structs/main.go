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
	//we could also use below just "contactInfo" and both name and type will be contactInfo
	contact contactInfo
}

func main() {
	// alex := person{
	// 	firstName: "Alex",
	// 	lastName:  "Anderson",
	// }

	var peter person
	peter.firstName = "Peter"
	peter.lastName = "Jones"

	//fmt.Println("Alex: ", alex)

	//This usage of printf allows us to print uninitialized properties.
	//Println would give '{}' in case we didn't specify any values
	//fmt.Printf("Peter: %+v", peter)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			zipCode: 20345,
			email:   "jim@email.com",
		},
	}

	fmt.Printf("Jim: %+v", jim)

	jim.print()

	jim.updateName("jimmy")

	jim.print()
}

//Pointers are back!
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
