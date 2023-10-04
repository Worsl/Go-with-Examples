package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)


type UserData struct {
    firstName string
	lastName string
	email string
	numberOfTickets int
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	for{

		if remainingTickets <= 0{
			fmt.Println("out of tickets")
			break
		}else{
			fmt.Println("Number of tickets remaining: ", remainingTickets)
		}
		
		firstName, lastName , email, userTickets := getUserInput()
		ValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName,lastName,email,userTickets)
	


		if ValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets,firstName,lastName,email)

			wg.Add(1)
			go sendTicket(userTickets,firstName,lastName,email)
		} else {
			if !isValidEmail{
				fmt.Println("invalid email")
			}
		
			if !ValidName{
				fmt.Println("invalid name")
			}
		
			if !isValidTicketNumber{
				fmt.Println("invalid ticket number")
			}
		}

		
	}
	
	
	fmt.Println("current list of users registered in the system")
	fmt.Println(getFirstNames())

}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}


func greetUsers() {
    fmt.Println("Welcome to the ticket booking conference session ")
	fmt.Printf("\n\n")
}

func getFirstNames() []string {
	s := make([]string,0)

    for _,b := range bookings{
		s = append(s, b.firstName + b.lastName )
	}

	return s
}

func getUserInput() (string, string, string, uint) {
	var userTickets uint
	var lastName string
	var firstName string
	var email string



    fmt.Print("Number of tickets to purchase: ")
	fmt.Scanln(&userTickets)

	fmt.Print("First name: ")
	fmt.Scanln(&firstName)

	fmt.Print("Last name: ")
	fmt.Scanln(&lastName)

	fmt.Print("email: ")
	fmt.Scanln(&email)


	return firstName,lastName,email,userTickets


}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	user := UserData{firstName: firstName, lastName: lastName, email: email, numberOfTickets: int(userTickets)}
	bookings = append(bookings, user)
	remainingTickets -= userTickets
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n%v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()

}