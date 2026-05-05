package main

import "fmt"

func checkLatestUser() {
	if len(users) == 0 {
		fmt.Println("No users available")
		return
	}

	user := users[len(users)-1]

	score := 0

	if len(user.FullName) > 10 {
		score++
	}

	if len(user.Email) > 10 {
		score++
	}

	if len(user.Phone) > 8 {
		score++
	}

	fmt.Println("\n--- MEMBERSHIP RESULT ---")

	if score >= 3 {
		fmt.Println("🟢 Gold Member")
	} else if score == 2 {
		fmt.Println("🟡 Silver Member")
	} else {
		fmt.Println("🔴 Bronze Member")
	}
}