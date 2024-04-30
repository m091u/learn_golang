package main

import "fmt"

// Define the ToyCar interface
type ToyCar interface {
	MoveForward()
	MoveBackward()
	MakeSound()
}

// Define a SportsCar struct
type SportsCar struct{}

// Implement the ToyCar interface for SportsCar
func (s SportsCar) MoveForward() {
	fmt.Println("Sports car moving forward")
}

func (s SportsCar) MoveBackward() {
	fmt.Println("Sports car moving backward")
}

func (s SportsCar) MakeSound() {
	fmt.Println("Sports car making sound")
}

func main() {
	// Create a SportsCar object
	sportsCar := SportsCar{}

	// Use the SportsCar object
	sportsCar.MoveForward()
	sportsCar.MoveBackward()
	sportsCar.MakeSound()
}
