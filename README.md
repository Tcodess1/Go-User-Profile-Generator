# Go User Profile Generator CLI

A beginner-friendly command-line application built with Go that collects a user’s name, generates a custom email address, analyses user details, and performs simple string operations using conditionals, loops, and slices.

This project was built as a practical coding challenge to reinforce core Go concepts through a complete mini-project that can be added to GitHub as part of a beginner developer portfolio.

---

## 📌 Project Overview

The **Go User Profile Generator CLI** is a simple command-line tool that allows a user to:

- Enter their first and last name
- Generate a custom email address
- Analyze their name
- compare phone size using conditional logic
- Manage and filter a list of countries
- Print a structured user profile summary

The goal of this project is to practice foundational Go programming concepts in a way that feels like building a real tool instead of isolated exercises.

---

## 🎯 What This Project Teaches

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

## ⚙️ Features

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
