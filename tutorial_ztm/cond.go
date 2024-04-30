package main

import "fmt"

const (
	Monday = 0
	Tuesday = 1
	Wednesday = 2
	Thursday = 3
	Friday = 4
	Saturday = 5
	Sunday = 6
)

const (
	Admin = 10
	Manager = 20
	Contractor = 30
	Member = 40
	Guest = 50
)

const (
	Economy = 0
	Business = 1
	FirstClass = 2
)

func price() int {
	return 1
}

func accessGranted(){
	fmt.Println("Granted")
}

func accessDenied(){
	fmt.Println("Denied")
}

func weekday (day int) bool {
	return day <=4
}