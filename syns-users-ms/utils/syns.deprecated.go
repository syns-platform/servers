package utils

import (
	"net/smtp"
	"os"
)

// @dev Report ip address to SYNS_PLATFORM_EMAIL
//
// @param ip string
//
// @notice DEPRECATED
func ReportVisitor(ip string) {
	// Set up authentication information for Gmail's SMTP server
	SYNS_EMAIL := os.Getenv("SYNS_PLATFORM_EMAIL")
	auth := smtp.PlainAuth("", SYNS_EMAIL, os.Getenv("SYNS_PLATFORM_EMAIL_PASSWORD"), "smtp.gmail.com")

	// Compose the email message
	to := []string{SYNS_EMAIL}
	msg := []byte("To: " +SYNS_EMAIL+ " +\r\n" +
		"Subject: Visitor IP address\r\n" +
		"\r\n" +
		"The IP address of the client is: " + ip + "\r\n")

	// Send the email using Gmail's SMTP server
	err := smtp.SendMail("smtp.gmail.com:587", auth, SYNS_EMAIL, to, msg)
	if err != nil {
		// Handle any errors that occur while sending the email
		panic(err)
	}
}