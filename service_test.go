package architecture

import (
	"fmt"
	"testing"
)

type Db map[int]Person

func (d Db) Save(n int, p Person) {
	d[n] = p
}

func (d Db) Retrieve(n int) Person {
	return d[n]
}

func TestPut(t *testing.T) {
	db := Db{}
	p := Person{
		First: "James",
	}

	Put(db, 1, p)

	if db.Retrieve(1) != p {
		t.Fatalf("Want %v, got %v", p, p)
	}
}

func ExamplePut() {
	db := Db{}
	p := Person{
		First: "James",
	}
	Put(db, 1, p)

	fmt.Println(db.Retrieve(1))
	// Output: {James}
}
