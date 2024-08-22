package main

import "strings"

func main() {

}

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	const message string = "Welcome to the Tech Palace"
	nameToUppCase := strings.ToUpper(customer)
	completMessage := message + ", " + nameToUppCase
	return completMessage

}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	star := strings.Repeat("*", numStarsPerLine)
	formatMessage := star + "\n" + welcomeMsg + "\n" + star
	return formatMessage
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msg := strings.ReplaceAll(oldMsg, "*", " ")
	msg = strings.TrimSpace(msg)
	return msg
}
