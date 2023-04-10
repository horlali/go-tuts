package main

import (
	"fmt"
	"strings"
)

func main() {
	var goTuts = "Go is awesome"
	const conferenceTickets = 50
	var remainigTickets uint = 50
	var bookings []string

	fmt.Println("Hello world")
	fmt.Println("Welcome to my go tutorials.", "I think", goTuts)
	fmt.Println("We have", conferenceTickets, "conference tickets and", remainigTickets, "remainig tickets")

	for {
		if remainigTickets == 0 {
			fmt.Println("Sorry we are sold out")
			break
		}
		var firstName string
		var lastName string
		var email string
		// var city string
		var numberOfTickets uint

		fmt.Println("Please enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email")
		fmt.Scan(&email)

		// fmt.Println("Please enter your city")
		// fmt.Scan(&city)

		fmt.Println("Please enter the number of tickets you want to book")
		fmt.Scan(&numberOfTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
		isValidNumberOfTickets := numberOfTickets > 0 && numberOfTickets <= remainigTickets
		// isValidCity := city == "Lagos" || city == "Abuja"

		if isValidEmail && isValidName && isValidNumberOfTickets {
			remainigTickets = remainigTickets - numberOfTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. We have sent a confirmation email to %v\n", firstName, lastName, numberOfTickets, email)
			fmt.Println("We have", remainigTickets, "remainig tickets")

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

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

		city := "Lagos"

		switch city {
		case "Lagos":
			fmt.Println("You are in Lagos")
		case "Abuja":
			fmt.Println("You are in Abuja")
		default:
			fmt.Println("You are not in Lagos or Abuja")
		}
	}

}
