package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceTickets int = 50  // const cannot be changed
	var remainingTickets uint = 50    // var can be changed
	conferenceName := "Go Conference" // shortcut for var definition (N/A for const)
	var bookings = []string{}         // empty slice of flexible size

	// print variable types
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
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

		remainingTickets = remainingTickets - userTickets

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

		firstNames := []string{} // shortcut for creating an empty slice
		// "_" is for unused param (index in this case)
		for _, booking := range bookings {
			// split the string with white space as separator, and return a slice with the split elements
			var names = strings.Fields(booking) // separate first name ans last name from a full name
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	}
}
