package main

import (
	"fmt"
)

type IBase interface {
	hello() string
}

type IMoreGreet interface {
	IBase
	goodbye() string
}

type Animal struct {
	Species string
}

type Person struct {
	FirstName string
	LastName  string
}

type BigPerson struct {
	Person
	Age int
}

type LazyPerson struct {
	IMoreGreet
}

func (p Person) hello() string {
	return "Hello, " + p.FirstName + " " + p.LastName
}

func (p Person) goodbye() string {
	return "Goodbye Mr " + p.LastName
}

func (a Animal) hello() string {
	return "Hi animal, " + a.Species
}

func (a Animal) goodbye() string {
	return "Get lost you cheeky " + a.Species
}

func sayStuff(base IMoreGreet) {
	fmt.Println(base.hello())
	fmt.Println(base.goodbye())
}

func main() {
	person := Person{"derek", "spink"}
	sayStuff(person)

	animal := Animal{"cat"}
	sayStuff(animal)

	bigPerson := BigPerson{Person{"clive", "tweebly"}, 50}
	//   bigPerson := BigPerson{FirstName: "reginald", LastName: "boo", Age: 60} // can't create a BigPerson like this

	fmt.Printf("We can also get the person out! %+v\n", bigPerson.Person)
	sayStuff(bigPerson)

	fmt.Println("oh no! " + bigPerson.FirstName) // but can refer to field of a bigPerson like this

	//lazyPerson := LazyPerson{}
	// sayStuff(lazyPerson) // compiles (as LazyPerson has IMoreGreet as anonymous interface) but panics at runtime as there is no implementation

}
