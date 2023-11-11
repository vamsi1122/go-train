package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var confName string = "KubeConf"

const confTickets int = 50

var remainTickets uint = 50
var booking = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, Email, userTickets := getUserInput()
		isValidEmail, isValidName, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, Email, userTickets, remainTickets)

		if isValidEmail && isValidName && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, Email)

			firstName := getFirstNames()

			fmt.Printf("The first names of booking are: %v\n", firstName)

			noTicketsavail := remainTickets == 0

			if noTicketsavail {
				//end
				fmt.Printf("Our conf is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or Last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email you entered is in valid")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickeets enetered is in valid")
			}
		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application and we have remaining tickets %v we begun at %v\n", confName, remainTickets, confTickets)
	fmt.Println("Get your tickets here")
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, bookings := range booking {
		var names = strings.Fields(bookings)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var Email string

	var userTickets int
	//asking user for the name
	fmt.Println("Enter User First Name")
	fmt.Scan(&firstName)
	fmt.Println("Enter User Last Name")
	fmt.Scan(&lastName)
	fmt.Println("Enter User Email")
	fmt.Scan(&Email)
	fmt.Println("Enter Tickets requireed")
	fmt.Scan(&userTickets)

	return firstName, lastName, Email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, Email string) {
	remainTickets = remainTickets - uint(userTickets)
	booking = append(booking, firstName+" "+lastName)

	fmt.Printf("User has booked %v under the name %v %v and the confirmation email will be sent to %v\n", userTickets, firstName, lastName, Email)

	fmt.Printf("%v tickets remaining for %v\n", remainTickets, confName)

}
