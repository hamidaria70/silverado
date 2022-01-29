package main

import (
	"fmt"
	"strings"
)

func main() {

	const conferenceTickets int = 50
	var conferenceName string = "Go Conference"
	var remainingTickets int = 50
	var bookings []string

	fmt.Printf("Welcome to our %v Booking App\n", conferenceName)
	fmt.Printf("Availabe tickets are %v and Remaining Tickets are %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var userName string
		var lastName string
		var userTickets int

		fmt.Println("Please enter your name:")
		fmt.Scan(&userName)
		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Please enter your tickets:")
		fmt.Scan(&userTickets)

		bookings = append(bookings, userName+" "+lastName)

		if userTickets <= remainingTickets{
			remainingTickets = remainingTickets - userTickets
			
			fmt.Printf("User %v %v booked %v tickets and remaining tickets are %v\n", userName, lastName, userTickets, remainingTickets)

			var firstNames []string
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Tickets are sold out")
				break
			}
		} else {
			fmt.Printf("%v is more than available(%v) tickets\n",userTickets, remainingTickets)
		}
		
	}
}
