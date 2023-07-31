package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference" // type string
	const conferenceTickets = 50         // constant value that never changes
	var remainingTickets uint = 50       // unsigned int (positive integers only)
	var bookings []string                // slice type

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We currently have %v of %v tickets available\n", remainingTickets, conferenceTickets)
	fmt.Println("Get you tickets here to attend...")

	// infinite loop
	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName) // get user input and store it in pointer (&variable)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Print("Enter the number of tickets you want to purchase: ")
		fmt.Scan(&userTickets)

		// append returns the updated slice
		bookings = append(bookings, firstName+" "+lastName)

		firstNames := []string{} // := type inference

		for _, booking := range bookings { // range returns an index (which we ignore), and item of slice at index
			var names = strings.Fields(booking) // split string on space, returns a slice
			firstNames = append(firstNames, names[0])
		}
		fmt.Println("This a list of those attending:", firstNames)

		remainingTickets = remainingTickets - userTickets
		fmt.Println("Info:", firstName, lastName, email, userTickets)
		fmt.Println(remainingTickets, "tickets remain")
	}
}
