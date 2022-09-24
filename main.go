package main

import "fmt"

func main() {
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	conferenceName := "Go Conference" // shortcut for var definition

	// print variable types
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask user input for their details
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName) // "&firstName" is a pointer, which is memory address like 0xc000014270

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Please enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
}
