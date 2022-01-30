package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	"sync"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)
var conferenceName string = "Go Conference"

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}


func main() {

	greetUser()

	var userName,lastName,userTickets,email = getUserInfo()
	var isValidEmail, isValidName, isValidTicketNumber = helper.InputValidation(userName,lastName,email,userTickets,remainingTickets)

	if isValidEmail && isValidName && isValidTicketNumber {
		remainingTickets = remainingTickets - userTickets

		var userData = UserData {
			firstName: userName,
			lastName: lastName,
			email: email,
			numberOfTickets: userTickets,
		}
		
		bookings = append(bookings, userData)

		fmt.Printf("the whole bookings is %v\n", bookings)
		fmt.Printf("User %v %v booked %v tickets and remaining tickets are %v\n", userName, lastName, userTickets, remainingTickets)
		wg.Add(1)
		go sendTickets(userTickets, userName, lastName, email)

		var firstNames []string = getFirstNames(bookings)
		fmt.Printf("These are all our bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Tickets are sold out")
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
	wg.Wait()
}

var wg = sync.WaitGroup{}

func greetUser() {
	fmt.Printf("Welcome to our %v Booking App\n", conferenceName)
	fmt.Printf("Availabe tickets are %v and Remaining Tickets are %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames([]UserData) []string{
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
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

func sendTickets(userTickets uint, userName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, userName, lastName)
	fmt.Println("******************************")
	fmt.Printf("Sending Ticket %v\n To email address %v\n", ticket, email)
	fmt.Println("******************************")
	wg.Done()
}