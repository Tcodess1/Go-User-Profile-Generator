# Golf Membership Checker CLI

A beginner-friendly command-line application built with Go that collects user information, generates a profile, and evaluates whether a user qualifies as a **Premium Golf Club Member**.

This project is designed to help beginners understand how real Go applications are structured using multiple files, while practicing core programming concepts like functions, structs, conditionals, and string manipulation.

---

## Project Overview

The **Go User Profile Generator CLI** is a simple command-line tool that allows a user to:

- Enter their first and last name
- Generate a custom email address
- Analyze their name
- compare phone size using conditional logic
- Manage and filter a list of countries
- Print a structured user profile summary

The goal is not just functionality, but learning how to structure a Go project properly.

---

## What This Project Teaches

This project helps reinforce the following Go concepts:

- Variables
- Data types
- String manipulation
- String slicing
- Indexing
- `strings.Split()`
- Arithmetic operations
- Boolean logic
- Conditional statements (`if`, `else if`, `else`)
- Arrays and slices
- Looping with `for`
- Iterating with `range`
- Basic filtering logic
- Output formatting with `fmt.Println`

---

## Features

### 1. User Name Collection
The program asks the user for their first and last names and combines both into a full name.

### 2. Email Generation
It automatically generates a custom email using:
- The first letter of the first name
- the full last name
- lowercase formatting
- `@01edu.net`

Example:

```text
Kesiah Valkin → kvalkin@01edu.net
