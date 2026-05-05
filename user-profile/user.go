package main

import (
	"fmt"
	"strings"
)

type User struct {
	FullName string
	Email    string
	Phone    string
}

var users []User

func createUserFlow() {
	var firstName, lastName, phone string

	fmt.Print("First name: ")
	fmt.Scanln(&firstName)

	fmt.Print("Last name: ")
	fmt.Scanln(&lastName)

	fmt.Print("Phone: ")
	fmt.Scanln(&phone)

	fullName := firstName + " " + lastName
	email := strings.ToLower(string(firstName[0]) + lastName + "@01edu.net")

	user := User{
		FullName: fullName,
		Email:    email,
		Phone:    phone,
	}

	users = append(users, user)

	fmt.Println("User created successfully!")
}

func showAllUsers() {
	if len(users) == 0 {
		fmt.Println("No users found.")
		return
	}

	for i, u := range users {
		fmt.Println("\nUser", i+1)
		fmt.Println("Name:", u.FullName)
		fmt.Println("Email:", u.Email)
		fmt.Println("Phone:", u.Phone)
	}
}