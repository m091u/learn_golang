package main

import (
	"strings"
)

func ValidateUserInput(userName string, lastName string, email string, userTickets uint, remainingTickets uint) (isValidName, isValidEmail, isValidTicketNumber bool) {
	isValidName = len(userName) >= 2 && len(lastName) >=2 
	isValidEmail = strings.Contains(email, "@")
	isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}


