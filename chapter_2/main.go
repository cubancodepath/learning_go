package main

import "fmt"

func main() {
	var i int = 20
	var f float64 = float64(i)
	fmt.Printf("i: %v\n", i)
	fmt.Println("f:", f)

	const value = 10

	i2, f2 := value, value

	fmt.Println("i2:", i2)
	fmt.Println("f2:", f2)

	var b int8 = 127
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
