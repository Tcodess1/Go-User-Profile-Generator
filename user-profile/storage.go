package main

import (
	"fmt"
	"os"
)

func saveUsersToFile() {
	file, err := os.Create("users.txt")
	if err != nil {
		fmt.Println("Error creating file")
		return
	}
	defer file.Close()

	for _, u := range users {
		line := u.FullName + " | " + u.Email + " | " + u.Phone + "\n"
		file.WriteString(line)
	}

	fmt.Println("Users saved to file")
}