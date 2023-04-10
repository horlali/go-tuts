package main

import (
	"fmt"
	"strings"
)

const conferenceName = "Ticket Booker World"
const conferenceTickets = 50

var remainigTickets uint = 50
var bookings = []string{}

func main() {
	greetUser()

	for {
		if remainigTickets == 0 {
			fmt.Println("Sorry we are sold out")
			break
		}

		firstName, lastName, email, numberOfTickets := getUserInput()
		isValidName, isValidEmail, isValidNumberOfTickets := validateUserInput(firstName, lastName, email, numberOfTickets)

		if isValidName && isValidEmail && isValidNumberOfTickets {
			bookTicket(firstName, lastName, email, numberOfTickets)

			firstNames := getFirstNames()
			fmt.Printf("Earliest person to book tickets: %v\n", firstNames[0])
			fmt.Printf("See other people who have booked tickets: %v\n", firstNames)

		} else {
			if !isValidName {
				fmt.Println("Name must be at least 2 characters")
			}
			if !isValidEmail {
				fmt.Println("Email must be valid")
			}
			if !isValidNumberOfTickets {
				fmt.Println("Sorry we only have", remainigTickets, "remainig tickets")
			}
		}

	}

}

func greetUser() {
	fmt.Printf("Hello\nWelcome to %v", conferenceName)
	fmt.Println("We have", conferenceTickets, "conference tickets and", remainigTickets, "remainig tickets")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, numberOfTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidNumberOfTickets := numberOfTickets > 0 && numberOfTickets <= remainigTickets

	return isValidName, isValidEmail, isValidNumberOfTickets
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var numberOfTickets uint

	fmt.Print("First Name: ")
	fmt.Scan(&firstName)

	fmt.Print("Last Name: ")
	fmt.Scan(&lastName)

	fmt.Print("Email: ")
	fmt.Scan(&email)

	fmt.Print("Number of tickets: ")
	fmt.Scan(&numberOfTickets)

	return firstName, lastName, email, numberOfTickets
}

func bookTicket(firstName string, lastName string, email string, numberOfTickets uint) {
	remainigTickets = remainigTickets - numberOfTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. We have sent a confirmation email to %v\n", firstName, lastName, numberOfTickets, email)
	fmt.Println("We have", remainigTickets, "remainig tickets")

}
