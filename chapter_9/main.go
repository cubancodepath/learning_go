package main

import (
	"errors"
	"fmt"
)

type InvalidId string

var InvalidIdError InvalidId = InvalidId("invalid id")

func (err InvalidId) Error() string {
	return "invalid id"
}

func (err InvalidId) Is(target error) bool {
	_, ok := target.(InvalidId)
	return ok
}

func createInvalidError() error {
	return InvalidId("invalid id")
}

type EmptyFieldErr struct {
	Field string
	Err   error
}

func (err EmptyFieldErr) Error() string {
	return fmt.Sprintf("field %s is empty", err.Field)
}

func (err EmptyFieldErr) Unwrap() error {
	return err.Err
}

type Employee struct {
	Name string
	Age  int
}

func (e *Employee) Validate() error {
	var errs = []error{}

	if e.Name == "" {
		errs = append(errs, EmptyFieldErr{Field: "name"})
	}
	if e.Age == 0 {
		errs = append(errs, EmptyFieldErr{Field: "age"})
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
func main() {
	err := createInvalidError()
	if errors.As(err, &InvalidIdError) {
		fmt.Println(err)
	}

	instance := Employee{Name: "Julian"}
	err = instance.Validate()
	if err != nil {
		if errors.As(err, &EmptyFieldErr{}) {
			fmt.Println(err)
		}
	}
}
