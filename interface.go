package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHelloI(hasName HasName) {
	fmt.Println("Hallo", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var Iqbal = Person{
		Name: "Iqbal",
	}

	sayHelloI(Iqbal)
}
