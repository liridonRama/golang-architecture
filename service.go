package architecture

import (
	"fmt"
)

// Person does stuff
type Person struct {
	First string
}

// Accessor is an interesting interface
type Accessor interface {
	Save(n int, p Person)
	Retrieve(n int) Person
}

// PersonService is nice
type PersonService struct {
	A Accessor
}

// Get gets stuff from db
func (ps PersonService) Get(n int) (Person, error) {
	p := ps.A.Retrieve(n)
	if p.First == "" {
		return Person{}, fmt.Errorf("No person with n of %d", n)
	}

	return p, nil
}

// Put adds stuff to the db
func Put(a Accessor, n int, p Person) {
	a.Save(n, p)
}

// Get gets it
func Get(a Accessor, n int) Person {
	return a.Retrieve(n)
}
