# Golf Membership Checker CLI

A beginner-friendly command-line application built with Go that collects user information, generates a profile, and evaluates whether a user qualifies as a **Premium Golf Club Member**.

This project is designed to help beginners understand how real Go applications are structured using multiple files, while practicing core programming concepts like functions, structs, conditionals, and string manipulation.

---

## Poject Overview

The Golf Club Membership Checker CLI simulates a simple membership system for a golf club.

It:

- collects user details from the terminal
- generates a custom email address
- analyses user profile data
- evaluates membership status using simple rules
- displays a final result in a structured format

The goal is not just functionality, but learning how to structure a Go project properly.

---

## What the Project Does

When you run the program, it:

### 1. Collects User Information

The program prompts the user to enter:

- First name
- Last name
- Phone number

---

### 2. Generates a User Profile

From the input, it creates:

- Full name
- Email address (first initial + last name + @01edu.net)

Example:

```text
Kesiah Valkin → kvalkin@01edu.net
```

---

### 3. Analyse User Data

The program also evaluates:

- Length of full name
- Basic string operations
- Simple data formatting

---

### 4. Membership Evaluation System

The system checks if a user qualifies as a premium member using a simple scoring logic:

### 🟢 Premium Member conditions

A user earns points based on:

- Valid email structure
- Strong phone number (numeric threshold check)
- Longer full name

### Result

- **2 or more points → 🟢 Premium Member**
- **Below 2 points → 🟡 Regular Member**

---

## Concepts You Will Learn

This project helps you practice:

### Go Basics

- Variables
- Data types
- Functions
- Structs
- Packages

### String Handling

- Concatenation
- Slicing
- Lowercasing
- Formatting

### Programming Logic

- If/else conditions
- Simple scoring system
- Decision making

### Data Structures

- Structs (`User`)
- Slices (country logic if extended)

### Software Design

- Multi-file project structure
- Separation of concerns
- Modular programming

---

## Project Structure

This project is split into multiple files to simulate real-world Go applications:

```text id="structure1"
go-user-profile-generator/
│── main.go          # Entry point (controls flow)
│── user.go          # User input & profile creation
│── membership.go    # Membership evaluation logic
│── utils.go         # Helper functions (reusable logic)
```

---

## How the Program Works

### Step 1: Program Starts

`main.go` runs the application.

---

### Step 2: User Input

`user.go` collects:

- First name  
- Last name  
- Phone number  

---

### Step 3: Profile Creation

A `User` struct is created containing:

- Full name  
- Generated email  
- Phone number  

---

### Step 4: Membership Check

`membership.go` evaluates the user using scoring rules.

---

### Step 5: Result Display

The final membership status is printed to the terminal.

---

## How to Run the Project

Make sure you are inside the project folder:

```bash
go run .
```

---

## Sample Output

```output
Welcome to Golf Club Membership Checker

Enter first name: Kesiah  
Enter last name: Valkin  
Enter phone number: 1234567890  

--- USER SUMMARY ---
Name: Kesiah Valkin  
Email: kvalkin@01edu.net  
Phone: 1234567890  

--- MEMBERSHIP CHECK ---
Status: 🟢 Premium Golf Club Member
```

---

## File Responsibilities

### main.go

- Starts the program  
- Controls application flow  

---

### user.go

- Handles user input  
- Builds user profile (struct)  
- Generates email  

---

### membership.go

- Contains business logic  
- Calculates membership score  
- Determines user status  

---

### utils.go

- Contains reusable helper functions  
- Supports clean code structure  

---

## Future Improvements

This project can be upgraded with:

- Multiple-user system  
- Saving user data to a file  
- Better email validation (regex)  
- Menu-driven CLI interface  
- Membership levels (Bronze, Silver, Gold, Premium)  
- Database integration  

---

## Learning Outcome

After completing this project, you will understand:

- How real Go projects are structured  
- How to split logic into multiple files  
- How structs work in real applications  
- How to design simple decision systems  
- How CLI applications are built  

---

## Author

Built as a learning project to improve:

- Go fundamentals  
- Clean code structure  
- Real-world project thinking  
