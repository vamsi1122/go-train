package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, Email string, userTickets int, remainTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	isValidEmail := strings.Contains(Email, "@")

	isValidTicketNumber := userTickets > 0 && userTickets <= int(remainTickets)

	return isValidName, isValidEmail, isValidTicketNumber
}
