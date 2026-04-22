package main

import "fmt"

func main() {
	greetings := []string{
		"Hello",
		"Hola",
		"नमस्ते",
		"こんにちは",
		"Привет",
	}

	firstTwo := greetings[:2]
	secondAndThird := greetings[1:3]
	forthAndFifth := greetings[3:]

	fmt.Println(firstTwo)
	fmt.Println(secondAndThird)
	fmt.Println(forthAndFifth)

	message := "Hi 👩🏼‍🦰 and 🧔🏻‍♂️"
	var messageRune []rune = []rune(message)

	for i := range messageRune {
		fmt.Println(string(messageRune[i]))
	}

	type Employee struct {
		firstName string
		lastName  string
		id        int
	}
	instance1 := Employee{
		"Bob",
		"Perez",
		12,
	}

	instance2 := Employee{
		firstName: "Kylian",
		lastName:  "Mabappe",
		id:        10,
	}

	var instance3 Employee

	instance3.firstName = "Vinicius"
	instance3.lastName = "Jr"
	instance3.id = 7

	fmt.Println(instance1)
	fmt.Println(instance2)
	fmt.Println(instance3)
}
