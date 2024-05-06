package main

import "fmt"

type contactInfo struct {
	address string
	zipCode int16
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) printDetails() {
	fmt.Printf("%+v", p)
}

func (p person) setFirstName(firstName string) {
	p.firstName = firstName
}

func main() {
	var om person
	om.printDetails()
	om = person{firstName: "om", lastName: "sirvi"}
	om.printDetails()
	om.setFirstName("omprakash")
	om.printDetails()
}
