package main

import (
	"fmt"
	"strings"
)

const totalTickets = 3

var availableTickets uint = 3

type guest = struct {
	firstName    string
	lastName     string
	email        string
	numOfTickets uint
}

var guests []guest

func getStrInput(prompt string) string {
	// get str input from user
	fmt.Print(prompt)
	var userInput string
	fmt.Scanln(&userInput)
	return userInput
}

func getIntInput(prompt string) uint {
	// get int input from user
	fmt.Print(prompt)
	var userInput uint
	fmt.Scanln(&userInput)
	return userInput
}

func welcomeMsg() {
	fmt.Println("\n\nWelcome to Go Conference booking app.")
	fmt.Printf("Book your tickets to attend. %v of %v tickets remaining...\n", availableTickets, totalTickets)
}

func isValidName(name string) (bool, string) {
	return (len(name) > 2), "Entered name is too short"
}

func isValidEmail(email string) (bool, string) {
	return (strings.Contains(email, "@")), "Invalid email address"
}

func validTickets(tickets uint) (bool, string) {
	if tickets > availableTickets {
		return false, fmt.Sprintf("Too many booked tickets. Only %v tickets remaining", availableTickets)
	} else {
		return true, ""
	}
}

func getNValidateName(prompt string) string {
	for {
		name := getStrInput(prompt)
		name_valid, reason := isValidName(name)
		if name_valid {
			return name
		} else {
			fmt.Printf("%v, try again.\n", reason)
		}
	}
}

func getNValidateEmail(prompt string) string {
	for {
		email := getStrInput(prompt)
		email_valid, reason := isValidEmail(email)
		if email_valid {
			return email
		} else {
			fmt.Printf("%v, try again.\n", reason)
		}
	}
}

func getNValidateTickets(prompt string) uint {
	for {
		ticketsBought := getIntInput(prompt)
		valid_tickets, reason := validTickets(ticketsBought)
		if valid_tickets {
			return ticketsBought
		} else {
			fmt.Printf("%v, try again.\n", reason)
		}
	}
}

func main() {
	// main function
	for {
		if availableTickets > 0 {
			var guestUser guest
			welcomeMsg() // print the welcome message and info

			// get and validate name from user before returning it
			first_name := getNValidateName("Enter your firstname: ")
			guestUser.firstName = first_name

			last_name := getNValidateName("Enter your lastname: ")
			guestUser.lastName = last_name

			// get email from user, validate and return it
			email := getNValidateEmail("Enter your email address: ")
			guestUser.email = email

			// get and validate number of tickets user booked
			ticketsBought := getNValidateTickets("Enter number of tickets to book: ")
			guestUser.numOfTickets = ticketsBought

			availableTickets -= ticketsBought // update count of available tickets
			guests = append(guests, guestUser)
			// display guests
			fmt.Println(guests)
		} else {
			fmt.Println("Tickets sold out. Thank you.")
			break
		}

	}
}
