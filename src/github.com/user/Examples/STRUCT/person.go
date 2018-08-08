package main

import "fmt"

type person struct {
	fname string
	lname string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFName(newFName string) {
	(*p).fname = newFName
}
