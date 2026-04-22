package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := []int{}
	for i := 0; i < 100; i++ {
		numbers = append(numbers, rand.Intn(100))
	}

	for value := range numbers {

		switch {
		case value%2 == 0 && value%3 == 0:
			fmt.Println("Six")
		case value%2 == 0:
			fmt.Println("Two")
		case value%3 == 0:
			fmt.Println("Three")
		default:
			fmt.Println("Never mind")
		}
	}

	var total int
	for range 10 {
		total := total + 1 //shadowing the original total variable we can fixed using total=total+1 or total++
		fmt.Println(total)
	}
}
