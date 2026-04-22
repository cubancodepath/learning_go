package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type operationFunType func(a, b int) (int, error)

var add operationFunType = func(a, b int) (int, error) {
	return a + b, nil
}
var sub = func(a, b int) (int, error) {
	return a - b, nil
}

var mul = func(a, b int) (int, error) {
	return a * b, nil
}

var div = func(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by 0")
	}
	return a - b, nil
}

var operationsMap = map[string]operationFunType{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func fileLen(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, errors.New("error open the file")
	}
	defer f.Close()
	b, err := f.Stat()
	if err != nil {
		return 0, errors.New("error reading file stats")
	}
	return int(b.Size()), nil
}

func prefixer(p string) func(value string) string {
	return func(value string) string {
		return p + " " + value
	}
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "3"},
		{"2", "/", "0"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Invalid expression")
			continue
		}
		op1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
		}
		op := expression[1]
		opeFunc, ok := operationsMap[op]
		if !ok {
			fmt.Println("Operator not suported")
			continue
		}
		op2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := opeFunc(op1, op2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(result)
	}

	filename := "./chapter_5"
	number, err := fileLen(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The number of bytes form \"%s\" is %v bytes\n", filename, number)

	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}
