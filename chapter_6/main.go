package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func makePerson(firstName, lastName string, age int) Person {
	return Person{
		LastName:  lastName,
		FirstName: firstName,
		Age:       age,
	}
}

func mekPersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		LastName:  lastName,
		FirstName: firstName,
		Age:       age,
	}
}

func UpdateSlice(slice []string, value string) {
	slice[len(slice)-1] = value
	fmt.Println(slice)
}

func GrowthSlice(slice []string, value string) {
	slice = append(slice, value)
	fmt.Println(slice)
}

func main() {
	person := makePerson("Bob", "Dowson", 45)
	personPointer := mekPersonPointer("John", "Doe", 35)

	fmt.Println(person)
	fmt.Println(personPointer)

	slice := []string{
		"Bob",
		"Jane",
		"Joe",
	}

	fmt.Println("Before Update", slice)
	UpdateSlice(slice, "John")
	fmt.Println("After Updated", slice)

	fmt.Println("-----")

	fmt.Println("Before Growth", slice)
	GrowthSlice(slice, "John")
	fmt.Println("After Growth", slice)

	persons := []Person{}
	for range 10_000_000 {
		persons = append(persons, Person{
			FirstName: "John",
			LastName:  "Doe",
			Age:       45,
		})
	}
}
