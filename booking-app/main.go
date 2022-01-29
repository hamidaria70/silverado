package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50
	var bookings []string

	greetUser(conferenceName, conferenceTickets, uint(remainingTickets))

	for {
		var userName string
		var lastName string
		var userTickets int
		var email string

		fmt.Println("Please enter your name:")
		fmt.Scan(&userName)
		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Please enter your email address:")
		fmt.Scan(&email)
		fmt.Println("Please enter your tickets:")
		fmt.Scan(&userTickets)

		var isValidName bool = len(userName) > 2 && len(lastName) > 2
		var isValidEmail bool = strings.Contains(email,"@")
		var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, userName+" "+lastName)

			fmt.Printf("User %v %v booked %v tickets and remaining tickets are %v\n", userName, lastName, userTickets, remainingTickets)

			printFirstNames(bookings)

			if remainingTickets == 0 {
				fmt.Println("Tickets are sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or Last name in too short")
			}
			if !isValidEmail {
				fmt.Println("Your email does not exist")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets is not valid")
			}
		}
	}
}

func greetUser(confName string, confTickets int, remainingTickets uint) {
	
	fmt.Printf("Welcome to our %v Booking App\n", confName)
	fmt.Printf("Availabe tickets are %v and Remaining Tickets are %v\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func printFirstNames(bookings []string){

	var firstNames []string
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("These are all our bookings: %v\n", firstNames)

}