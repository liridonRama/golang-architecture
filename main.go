package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I'm a person. this is my name", p.first)
}

func (s secretAgent) speak() {
	fmt.Println("I'm a secret agent — this is my name: ", s.first)
}

type human interface {
	speak()
}

func main() {
	p1 := person{
		first: "Miss Moneypenny",
	}
	s1 := secretAgent{person: p1, ltk: true}

	var x, y human
	x, y = p1, s1

	x.speak()
	y.speak()
}
