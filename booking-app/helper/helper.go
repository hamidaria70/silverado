package helper

import "strings"


func InputValidation(userName string,lastName string,email string,userTickets uint, RemainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(userName) > 2 && len(lastName) > 2
	var isValidEmail bool = strings.Contains(email,"@")
	var isValidTicketNumber bool = userTickets > 0 && userTickets <= RemainingTickets

	return isValidEmail, isValidName, isValidTicketNumber
}