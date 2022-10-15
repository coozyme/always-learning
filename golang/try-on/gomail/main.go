package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"

	"text/template"

	gomail "gopkg.in/gomail.v2"
)

type Mail struct {
	Name  string
	Email string
}

func (ml *Mail) sendMail() {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "activation@xxx.com")

	// Set E-Mail receivers
	m.SetHeader("To", "xxx@gmail.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "Activation")
	t := template.New("template.html")

	t, err := template.ParseFiles("template.html")

	if err != nil {
		fmt.Println(err)
	}

	var body bytes.Buffer

	if err := t.Execute(&body, ml); err != nil {
		log.Println(err)
	}
	// Set E-Mail body. You can set plain text or html with text/html
	// m.SetBody("text/plain", "This is Gomail test body")
	// result, _ := ParseTemplate("template.html", "programmeranak@gmail.com")
	result := body.String()
	m.SetBody("text/plain", result)

	// Settings for SMTP server
	d := gomail.NewDialer("mail.xxx.com", 587, "activation@xxx.com", "xxx")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err)
		return "", err
	}
	return buf.String(), nil
}

func main() {
	d := Mail{"jack", "xxx@gmai.com"}

	d.sendMail()

}
