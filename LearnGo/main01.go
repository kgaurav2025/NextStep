package main

import (
	"fmt"
)


func main() {
	var a int= 10
	var name string = "Gopher"
	fmt.Println("Value of a:", a, "Name:", name)
	fmt.Println("Hello, Go!")


	var age int = 25

	// Using var (type inferred)
	var city = "New York"

	// Short declaration (most common inside functions)
	count := 10
	message := "Hello"

	// Constants (cannot be changed)
	const Pi = 3.14159
	const MaxUsers = 100

	fmt.Println("Value of age:", age, "city:", city)
	fmt.Println("Value of count:", count, "count:", count)
	fmt.Println("Value of message:", message, "city:", city)
}