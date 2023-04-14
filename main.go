package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceName = "Ticket Booker World"
const conferenceTickets = 50

var remainigTickets uint = 50
var bookings = make([]UserData, 0)

type UserData = struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUser()

	if remainigTickets == 0 {
		fmt.Println("Sorry we are sold out")
		// break
	}

	firstName, lastName, email, numberOfTickets := getUserInput()
	isValidName, isValidEmail, isValidNumberOfTickets := validateUserInput(firstName, lastName, email, numberOfTickets, remainigTickets)

	if isValidName && isValidEmail && isValidNumberOfTickets {
		bookTicket(firstName, lastName, email, numberOfTickets)

		wg.Add(1)
		go sendTicket(firstName, lastName, email, numberOfTickets)

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
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Hello\nWelcome to %v", conferenceName)
	fmt.Println("We have", conferenceTickets, "conference tickets and", remainigTickets, "remainig tickets")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: numberOfTickets,
	}

	bookings = append(bookings, userData)
	fmt.Println(bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. We have sent a confirmation email to %v\n", firstName, lastName, numberOfTickets, email)
	fmt.Println("We have", remainigTickets, "remainig tickets")

}

func sendTicket(firstName string, lastName string, email string, numberOfTickets uint) {
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", numberOfTickets, firstName, lastName)
	fmt.Println("#############################################")
	fmt.Printf("Sending ticket:\n %vto email address %v\n", ticket, email)
	fmt.Println("#############################################")
	wg.Done()
}

// A few more things to learn
// 1. Maps in golang
// 2. Structs in golang
// 3. Coroutines and concurrency in golang
// 4. Interfaces in golang --  not in Nana's tutorial
// 5. How to write tests in golang --  not in Nana's tutorial
// 6. How to write benchmarks in golang --  not in Nana's tutorial
