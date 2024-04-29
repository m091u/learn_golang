package main

import (
	"fmt" //used for input /output
	// "strconv"
	// "strings"
	// "booking-app/helper"
	"time"
	"sync"
)

var confName = "Go Conference"
const confTickets uint = 50  
var remainingTickets uint = 50
var bookings = make([]UserData,0)

type UserData struct {
	userName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		userName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber :=  ValidateUserInput(userName,lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, userName, lastName, email)
			wg.Add(1)

			go sendTicket(userTickets, userName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			fmt.Println(bookings)
	
			if remainingTickets == 0{
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			} 
		} else {
			if !isValidName {
			    fmt.Println("First name or last name you entered is too short.")
			} 
			if !isValidEmail {
			    fmt.Println("Email address you entered doesn't contain @ sign.")
			}
			if !isValidTicketNumber {
			    fmt.Println("Number of tickets you entered is invalid.")
			}
		fmt.Println("Please try again.")
		}
	}
	wg.Wait()
}

func greetUsers(){
	fmt.Printf("Welcome to our %v booking app!\n", confName)
	fmt.Printf("We have a total of %v tickets available and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets for", confName, "here", )
}

func getFirstNames() []string{
	firstNames :=[]string{}
	for _, booking := range bookings{
		firstNames = append(firstNames, booking.userName)
	}
	return firstNames
}


func getUserInput()(string, string, string, uint){
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

	return userName, lastName, email, userTickets
}

func bookTicket( userTickets uint, userName string, lastName string, email string){
	remainingTickets = remainingTickets - userTickets

	// map for user

	var userData = UserData{
		userName: userName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you, %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confName)
}

func sendTicket(userTickets uint, userName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v this many ticvets for %v %v", userTickets, userName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")

	wg.Done()
} 