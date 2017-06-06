package main

import (
	"fmt"
)

// see https://stackoverflow.com/questions/24537443/meaning-of-a-struct-with-embedded-anonymous-interface for interesting discussion around aspects of this

type IGreet interface {
	hello() string
}

type ISalutation interface {
	IGreet
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

type CoughingPerson struct {
	ISalutation
}

func (coughingPerson CoughingPerson) hello() string {
	return "*cough* " + coughingPerson.ISalutation.hello() + " *cough*"
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

func giveCough(salutationer ISalutation) ISalutation {
	return &CoughingPerson{salutationer}
}

func sayStuff(base ISalutation) {
	fmt.Println(base.hello())
	fmt.Println(base.goodbye())
}

func main() {

	person := Person{"derek", "spink"}
	sayStuff(person) // runs the Person methods for hello/goodbye

	fmt.Println("---------------------------")

	animal := Animal{"cat"}
	sayStuff(animal) // runs the Animal methods for hello/goodbye

	fmt.Println("---------------------------")

	bigPerson := BigPerson{Person{"clive", "tweebly"}, 50}
	//   bigPerson := BigPerson{FirstName: "reginald", LastName: "boo", Age: 60} // can't create a BigPerson like this
	sayStuff(bigPerson)                          // since BigPerson is also a Person, will just run the Person hello/goodbye methods
	fmt.Println("oh no! " + bigPerson.FirstName) // but can refer to field of a bigPerson like this
	fmt.Printf("We can also get the person out! %+v\n", bigPerson.Person)

	fmt.Println("---------------------------")

	bigCoughingPerson := giveCough(bigPerson)
	sayStuff(bigCoughingPerson) // now uses the CoughingPerson method for hello(), which adds some coughs

	fmt.Println("---------------------------")

	coughingSomething := CoughingPerson{}
	// sayStuff(coughingSomething) // compiles (as CoughingPerson has ISalutation as anonymous interface) but panics at runtime as there is no implementation for ISalutation cos coughingSomething was not instantiated with an ISalutation instance
	coughingSomething.ISalutation = animal
	sayStuff(coughingSomething) // now is fine as we've given it an ISalutation instance.

}
