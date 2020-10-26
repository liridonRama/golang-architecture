package main

import "fmt"

type person struct {
	first string
}

type mongo map[int]person
type postgres map[int]person

func (m mongo) save(n int, p person) {
	m[n] = p
}

func (m mongo) retrieve(n int) person {
	return m[n]
}

func (pg postgres) save(n int, p person) {
	pg[n] = p
}

func (pg postgres) retrieve(n int) person {
	return pg[n]
}

type accessor interface {
	save(n int, p person)
	retrieve(n int) person
}

func put(a accessor, n int, p person) {
	a.save(n, p)
}

func get(a accessor, n int) person {
	return a.retrieve(n)
}

func main() {
	dbm := mongo{}
	dbpg := postgres{}

	p1 := person{"jenny"}
	p2 := person{"james"}
	p3 := person{"jonny"}

	put(dbm, 1, p1)
	put(dbm, 2, p2)
	put(dbm, 3, p3)
	put(dbpg, 1, p1)
	put(dbpg, 2, p2)
	put(dbpg, 3, p3)

	fmt.Println(get(dbm, 1).first)
	fmt.Println(get(dbpg, 2).first)
	fmt.Println(get(dbm, 3).first)

}
