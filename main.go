package main

import (
	"fmt"
	"math"
)

// Struct definition
type Point struct {
	X, Y int
}

// Function with parameters and return value
func add(a, b int) int {
	return a + b
}

// Function with multiple return values
func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return dividend / divisor, nil
}

// Interface definition
type Shape interface {
	Area() float64
}

// Struct implementing an interface
type Circle struct {
	Radius float64
}

// Method for the Circle struct to implement the Shape interface
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	// Variable declaration and assignment
	var x int = 10
	y := 5 // Type inference

	// Conditional statement
	if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("y is greater than or equal to x")
	}

	// Looping construct
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Array declaration and initialization
	numbers := [3]int{1, 2, 3}

	// Slice creation
	slice := numbers[1:3]

	// Map declaration and initialization
	person := map[string]string{
		"name":  "John",
		"age":   "30",
		"city":  "New York",
		"email": "john@example.com",
	}
	
	// Accessing values
	name := person["name"]
	age := person["age"]
	city := person["city"]
	email := person["email"]

	// Printing values
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %s\n", age)
	fmt.Printf("City: %s\n", city)
	fmt.Printf("Email: %s\n", email)

	// Modifying values
	person["age"] = "31"
	person["city"] = "San Francisco"

	// Printing updated values
	fmt.Printf("Updated Age: %s\n", person["age"])
	fmt.Printf("Updated City: %s\n", person["city"])
	

	// Calling a function
	sum := add(x, y)

	// Error handling
	result, err := divide(10.0, 2.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Using a struct
	p := Point{X: 1, Y: 2}

	// Printing values
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Point: %+v\n", p)

	// Looping through a slice
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Using an interface
	var shape Shape
	shape = Circle{Radius: 5.0}
	area := shape.Area()
	fmt.Printf("Circle Area: %.2f\n", area)
}
