package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

// package-level variables, shared across files using "package main" at the top
const conferenceTickets int = 50            // const cannot be changed
var remainingTickets uint = 50              // var can be changed
var conferenceName string = "Go Conference" // cannot use shortcut declaration (a := b) in package-level variables
var bookings = []string{}                   // empty slice of flexible size

func main() {

	greetUsers()

	// print variable types
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// indefinite loop, each loop for each user
	for {
		// go functions can return multiple values
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// // if email is invalid, use "continue" to skip the remainder of current loop, retry booking with another loop
		// if !isValidEmail {
		// 	fmt.Println("The email you entered does not contain @ sign, please try again")
		// 	continue
		// }

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			// firstNames is a slice of strings
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program (loop)
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("The email you entered does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers() {
	// using package-level variables directly, without parameters
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{} // shortcut for creating an empty slice
	// loop through bookings slice, "_" is for unused param (index in this case)
	for _, booking := range bookings {
		// split the string with white space as separator, and return a slice with the split elements
		var names = strings.Fields(booking) // separate first name ans last name from a full name by white space
		firstNames = append(firstNames, names[0])
	}
	// return a slice of strings
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user input for their details
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName) // "&firstName" is a pointer, which is memory address like 0xc000014270

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Please enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("bookings: %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
