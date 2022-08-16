package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Configuration
	from := "programmeranak@gmail.com"
	password := "r9$7@4T*5Nh8c7F2S*"
	to := []string{"appststng@gmail.com"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("My super secret message.")

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}
}
