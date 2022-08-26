package models

type Person struct {
	ID   int
	Name string
	Age  int
	Type int
}

const (
	user = iota + 1
	admin
	reader
	writer
)

type Persons []Person

func GetPersons() Persons {
	persons := make(Persons, 0, 5)
	persons = append(persons,
		Person{ID: 1, Name: "Chalchiuhcuauhtli", Age: 40, Type: admin},
		Person{ID: 2, Name: "User 1", Age: 10, Type: user},
		Person{ID: 3, Name: "Reader", Age: 20, Type: reader},
		Person{ID: 4, Name: "Writer", Age: 30, Type: writer},
		Person{ID: 5, Name: "User 2", Age: 17, Type: user},
		Person{ID: 6, Name: "Chalchiuhcuauhtli 2", Age: 35, Type: admin},
		Person{ID: 7, Name: "User 3", Age: 37, Type: user},
		Person{ID: 8, Name: "Reader 2", Age: 21, Type: reader},
		Person{ID: 9, Name: "Writer 2", Age: 19, Type: writer},
		Person{ID: 10, Name: "User 4", Age: 13, Type: user},
	)

	return persons
}
