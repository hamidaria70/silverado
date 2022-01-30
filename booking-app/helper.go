package main

import "strings"


func inputValidation(userName string,lastName string,email string,userTickets uint) (bool, bool, bool) {
	var isValidName bool = len(userName) > 2 && len(lastName) > 2
	var isValidEmail bool = strings.Contains(email,"@")
	var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTickets

	return isValidEmail, isValidName, isValidTicketNumber
}