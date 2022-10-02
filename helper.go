package main

import (
	"strings"
)

// use Capital func name to export it
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	// remainingTickets is package-level variable
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	// golang can return multiple values
	return isValidName, isValidEmail, isValidTicketNumber
}
