package main

import "fmt"

type contactInfo struct {
	emailAddress string
	zipCode      int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}

func main() {
	mike := person{
		firstName: "Mike",
		lastName:  "Smith",
		contactInfo: contactInfo{
			emailAddress: "mail@mikegsmith.co.uk",
			zipCode:      90210,
		},
	}

	mike.updateName("Michael")
	mike.print()
}
