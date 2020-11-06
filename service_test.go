package architecture

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Db map[int]Person

func (d Db) Save(n int, p Person) {
	d[n] = p
}

func (d Db) Retrieve(n int) Person {
	return d[n]
}

func TestPut(t *testing.T) {
	assert := assert.New(t)
	db := Db{}
	p := Person{
		First: "James",
	}

	p2 := Person{
		First: "hans",
	}

	Put(db, 1, p)

	assert.Equal(db.Retrieve(1), p, "should be the same person")
	assert.NotEqual(db.Retrieve(1), p2, "should not be equal")
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
