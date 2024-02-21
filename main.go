package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T \n", remainingTickets, conferenceTickets, conferenceName)

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask fr teir name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("The whole slice: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("slice type: %T\n", bookings)
			// fmt.Printf("slice length: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for the Go conference.\n", remainingTickets)

			// slice for only first names
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all the first names: %v\n", firstNames)
			// fmt.Printf("These are all our bookings: %v\n", bookings)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email you entered doesnt contain @")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invlaid")

			}
			fmt.Println("Your input data is invalid, try again")
			// fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets\n", remainingTickets, userTickets)
		}
	}
}
