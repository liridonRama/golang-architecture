package main

import (
	"fmt"

	architecture "github.com/liridonRama/golang-architecture"
	"github.com/liridonRama/golang-architecture/storage/mongo"
	"github.com/liridonRama/golang-architecture/storage/postgres"
)

func main() {
	dbm := mongo.Db{}
	dbpg := postgres.Db{}
	psmongo := architecture.PersonService{A: dbm}

	p1 := architecture.Person{First: "jenny"}
	p2 := architecture.Person{First: "james"}
	p3 := architecture.Person{First: "jonny"}

	architecture.Put(dbm, 1, p1)
	architecture.Put(dbm, 2, p2)
	architecture.Put(dbm, 3, p3)
	architecture.Put(dbpg, 1, p1)
	architecture.Put(dbpg, 2, p2)
	architecture.Put(dbpg, 3, p3)

	fmt.Println(architecture.Get(dbm, 1).First)
	fmt.Println(architecture.Get(dbpg, 2).First)
	fmt.Println(architecture.Get(dbm, 3).First)

	rp1, err := psmongo.Get(1)
	fmt.Println(rp1.First, err)

	rp2, err := psmongo.Get(5)
	fmt.Println(rp2.First, err)

}
