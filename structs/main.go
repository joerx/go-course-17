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
	staffNo   int
	contactInfo
}

func main() {
	// what if order of fields changes?
	alex := person{"Alex", "Anderson", 20, contactInfo{}}

	// more gud
	john := person{
		firstName: "John",
		lastName:  "Wick",
		contactInfo: contactInfo{
			zipCode: 01234,
			email:   "john.wick@continental.biz",
		},
	}

	var angela person
	angela.firstName = "Angela"
	angela.contactInfo.email = "angela@bundesregierung.gov"

	fmt.Println(alex)
	fmt.Println(john)
	fmt.Printf("%+v\n", angela)

	fmt.Println("---")

	alex.print()

	fmt.Println("---")

	alex.setFirstName("Alexandra")
	alex.print()

	fmt.Println("---")

	pAlex := &alex
	pAlex.firstName = "Mr"
	pointAt(pAlex)
}

func (p *person) setFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func pointAt(p *person) {
	fmt.Printf("Pointing at %s %s\n", p.firstName, p.lastName)
}

func (p person) print() {
	fmt.Printf("First name: %s\n", p.firstName)
	fmt.Printf("Last name: %s\n", p.lastName)
	fmt.Printf("Email: %s\n", p.email)
	fmt.Printf("Zipcode: %05d\n", p.zipCode)
}
