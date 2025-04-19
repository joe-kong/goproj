package main

import "fmt"

// Package main demonstrates a simple user management system
// and email service using structs and methods in Go.
// This program defines a User struct, a UserManager for user operations,
// and an EmailService to send emails.
// It showcases how to create and manipulate user data,
// and how to send a welcome email to a new user.
// It is a basic example of struct usage and method definitions in Go.　単一責任原則
// User represents a user in the system
type User struct {
	Name  string
	Email string
}

// UserManager handles user-related operations
type UserManager struct{}

func (um *UserManager) AddUser(user User) {
	fmt.Printf("User %s added successfully.\n", user.Name)
}

// EmailService handles email-related operations
type EmailService struct{}

func (es *EmailService) SendWelcomeEmail(user User) {
	fmt.Printf("Welcome email sent to %s.\n", user.Email)
}

func main() {
	user := User{Name: "John Doe", Email: "john.doe@example.com"}

	userManager := UserManager{}
	emailService := EmailService{}

	// Adding a user
	userManager.AddUser(user)

	// Sending a welcome email
	emailService.SendWelcomeEmail(user)
}
