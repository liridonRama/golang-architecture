package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person person
	ltk    bool
}

func (p person) speak() {
	fmt.Println("from a person", p.first)
}

func (s secretAgent) speak() {
	fmt.Println("from a person", s.person.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{"jenny"}
	sa1 := secretAgent{
		person: person{
			first: "james",
		},
		ltk: true,
	}

	saySomething(p1)
	saySomething(sa1)
}
