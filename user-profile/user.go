package main

import (
	"fmt"
	"strings"
)

type User struct {
	FullName string
	Email    string
	Phone    float64
}

func getUserProfile() User {
	var firstName string
	var lastName string
	var phone float64

	fmt.Print("Enter first name: ")
	fmt.Scanln(&firstName)

	fmt.Print("Enter last name: ")
	fmt.Scanln(&lastName)

	fmt.Print("Enter phone number (numeric): ")
	fmt.Scanln(&phone)

	fullName := firstName + " " + lastName
	firstInitial := string(firstName[0])

	email := strings.ToLower(firstInitial + lastName) + "@01edu.net"

	return User{
		FullName: fullName,
		Email:    email,
		Phone:    phone,
	}
}