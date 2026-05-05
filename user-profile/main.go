package main

import "fmt"

func main() {
	fmt.Println("🏌️ Welcome to Lagos Golf Club Membership Checker")

	user := getUserProfile()
	checkMembership(user)
}