package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, numberOfTickets uint, remainigTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidNumberOfTickets := numberOfTickets > 0 && numberOfTickets <= remainigTickets

	return isValidName, isValidEmail, isValidNumberOfTickets
}
