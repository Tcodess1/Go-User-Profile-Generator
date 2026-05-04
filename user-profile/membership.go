package main

import "fmt"

func checkMembership(user User) {
	fmt.Println("\n--- USER SUMMARY ---")
	fmt.Println("Name:", user.FullName)
	fmt.Println("Email:", user.Email)
	fmt.Println("Phone:", user.Phone)

	fmt.Println("\n--- MEMBERSHIP CHECK ---")

	score := 0

	// Rule 1: email check (basic validation)
	if len(user.Email) > 10 {
		score++
	}

	// Rule 2: phone strength
	if user.Phone > 1000000000 {
		score++
	}

	// Rule 3: name strength
	if len(user.FullName) > 10 {
		score++
	}

	// Decision
	if score >= 2 {
		fmt.Println("Status: 🟢 Premium Golf Club Member")
	} else {
		fmt.Println("Status: 🟡 Regular Member")
	}
}