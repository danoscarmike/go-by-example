package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	var p person
	p = person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(newPerson("Jon"))
	dan := person{name: "Dan", age: 42}
	fmt.Println(dan)
	hugh := newPerson("Hugh")
	hugh.age = 5
	fmt.Println(hugh)
	dan.age = 43
	fmt.Println(dan)
}
