package main

import (
	"fmt"
)

func main() {

	const conferenceTickets int = 50
	var conferenceName string = "Go Conference"
	var remainingTickets int = 50
	var userName string
	var userTickets int

	fmt.Printf("Welcome to our %v Booking App\n", conferenceName)
	fmt.Printf("Availabe tickets are %v and Remaining Tickets are %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("Please enter your name:")

	fmt.Scan(&userName)
	fmt.Println("Please enter your tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v booked %v tickets and remaining tickets are %v\n", userName, userTickets, remainingTickets)
}
