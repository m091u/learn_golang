package main

import (
	"fmt" //used for input /output 
	"strings"
)

func main() {
	confName := "Go Conference"
	const confTickets = 50  
	var remainingTickets uint = 50
	var bookings []string

    fmt.Printf("Welcome to our %v booking app!\n", confName)
	fmt.Printf("We have a total of %v tickets available and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets for", confName, "here", )

	for i:=0; i<=50; i++ {
		var userName string
		var lastName string
		var email string
		var userTickets uint
	
		fmt.Println("Enter your name: ")
		//Scan() used to get the value from the user. Parameter is the input it needs to scan for
		fmt.Scan(&userName) 
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
	
		fmt.Printf("%v, how many tickets would you like to book?\n", userName)
		fmt.Scan(&userTickets)
	
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, userName +" "+ lastName)
	
		fmt.Printf("Thank you, %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confName)
	
		firstNames :=[]string{}
		for _, booking := range bookings{
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	}

}