package main

import (
	"fmt"
	"strings"
	"booking-app/helper"
)

const conferenceTickets int = 50
var bookings []string
var remainingTickets uint = 50
var conferenceName string = "Go Conference"

func main() {

	greetUser()

	for {
		var userName,lastName,userTickets,email = getUserInfo()
		var isValidEmail, isValidName, isValidTicketNumber = helper.InputValidation(userName,lastName,email,userTickets,remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			
			var bookings []string = bookTicket(userTickets, userName, lastName)


			var firstNames []string = getFirstNames(bookings)
			fmt.Printf("These are all our bookings: %v\n", firstNames)

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

func greetUser() {
	fmt.Printf("Welcome to our %v Booking App\n", conferenceName)
	fmt.Printf("Availabe tickets are %v and Remaining Tickets are %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string{
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInfo() (string, string, uint, string) {
	var userName string
	var lastName string
	var userTickets uint
	var email string

	fmt.Println("Please enter your name:")
	fmt.Scan(&userName)
	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Please enter your tickets:")
	fmt.Scan(&userTickets)

	return userName,lastName, userTickets,email
}

func bookTicket(userTickets uint , userName string, lastName string) []string{
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userName+" "+lastName)

	fmt.Printf("User %v %v booked %v tickets and remaining tickets are %v\n", userName, lastName, userTickets, remainingTickets)
	return bookings
}