package main

import (
	"fmt"
)

func main() {

	const conferenceTickets int = 50
	var conferenceName string = "Go Conference"
	var remainingTickets int = 50
	var userName string
	var lastName string
	var userTickets int
	var bookings []string

	fmt.Printf("Welcome to our %v Booking App\n", conferenceName)
	fmt.Printf("Availabe tickets are %v and Remaining Tickets are %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("Please enter your name:")
	fmt.Scan(&userName)
	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your tickets:")
	fmt.Scan(&userTickets)

	bookings = append(bookings, userName+" "+lastName)

	fmt.Printf("first element of slice is %v\n", bookings[0])
	fmt.Printf("the whole slice is %v\n", bookings)
	fmt.Printf("The type of slice is %T\n", bookings)
	fmt.Printf("The length of slice is %v\n", len(bookings))
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v %v booked %v tickets and remaining tickets are %v\n", userName, lastName, userTickets, remainingTickets)
}
