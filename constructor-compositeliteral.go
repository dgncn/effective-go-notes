package main

import "fmt"

func main() {
	//constructor example 1
	person1 := constructorFunc(true, 0)
	person2 := constructorFunc(false, 81)
	fmt.Printf("1: %v 2: %v", person1, person2)

	//composite literal example 2
	compositeLiteralFunc(true, 0)
	compositeLiteralFunc(false, 81)
}

func constructorFunc(isAlive bool, deathYear int8) *Person {
	var p *Person = new(Person)
	p.IsAlive = isAlive
	if p.IsAlive == true {
		p.DeathYear = 0
	} else {
		p.DeathYear = deathYear
	}
	return p
}

func compositeLiteralFunc(isAlive bool, deathYear int8) (person *Person) {
	person = &Person{
		IsAlive:   isAlive,
		DeathYear: deathYear,
	}
	fmt.Printf("value: %v pointer address: %v", person, &person)

	//new(Person) ifadesi &Person{} a eşdeğerdir.
	return
}

type Person struct {
	IsAlive   bool
	DeathYear int8
}
