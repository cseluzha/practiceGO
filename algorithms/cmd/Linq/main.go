package main

import (
	"algorithms/internal/models"
	"fmt"
	l "github.com/ahmetalpbalkan/go-linq"
)

func main() {
	persons := models.GetPersons()
	// TODO:
	// LINQ with type of data
	// Sort
	// Count
	// Iterator

	var minors []string

	l.From(persons).Where(func(p interface{}) bool {
		return p.(models.Person).Age < 15
	}).Select(func(p interface{}) interface{} {
		return p.(models.Person).Name
	}).ToSlice(&minors)

	fmt.Println(minors)

	collection := l.From(persons).OrderByDescendingT(func(p models.Person) int {
		return p.Age
	})

	fmt.Println(collection.Results()...)
}
