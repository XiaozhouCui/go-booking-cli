package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceTickets int = 50  // const cannot be changed
	var remainingTickets uint = 50    // var can be changed
	conferenceName := "Go Conference" // shortcut for var definition with type inference (N/A for const)
	var bookings = []string{}         // empty slice of flexible size

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// print variable types
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// indefinite loop, each loop for each user
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

		// go func can return multi values
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// // if email is invalid, use "continue" to skip the remainder of current loop, retry booking with another loop
		// if !isValidEmail {
		// 	fmt.Println("The email you entered does not contain @ sign, please try again")
		// 	continue
		// }

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

			// firstNames is a slice of strings
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program (loop)
				fmt.Println("Our conference is booked out. Come back next year.")
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
				fmt.Printf("We only have %v tickets remaning, so you can't book %v tickets\n", remainingTickets, userTickets)
			}
		}
	}
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
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

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	// golang can return multiple values
	return isValidName, isValidEmail, isValidTicketNumber
}
