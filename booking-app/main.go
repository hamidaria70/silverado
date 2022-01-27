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

	fmt.Printf("Welcome to our %v Booking App\n", conferenceName)
	fmt.Printf("Availabe tickets are %v and Remaining Tickets are %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("Please enter your name:")
	fmt.Scan(&userName)
	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your tickets:")
	fmt.Scan(&userTickets)

	var bookings [50]string
	bookings[0] = userName + " " + lastName

	fmt.Printf("first element of array is %v\n", bookings[0])
	fmt.Printf("the whole array is %v\n", bookings)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v %v booked %v tickets and remaining tickets are %v\n", userName, lastName, userTickets, remainingTickets)
}
