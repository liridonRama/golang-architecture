package postgres

import architecture "github.com/liridonRama/golang-architecture"

type Db map[int]architecture.Person

func (d Db) Save(n int, p architecture.Person) {
	d[n] = p
}

func (d Db) Retrieve(n int) architecture.Person {
	return d[n]
}
