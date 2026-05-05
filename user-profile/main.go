package main

import "fmt"

func main() {
	for {
		showMenu()

		var choice int
		fmt.Print("\nEnter choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			createUserFlow()
		case 2:
			showAllUsers()
		case 3:
			checkLatestUser()
		case 4:
			fmt.Println("Goodbye 👋")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}