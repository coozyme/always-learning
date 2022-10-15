package main

import (
	"bytes"
	"fmt"
	"html/template"
	"runtime"

	"gopkg.in/gomail.v2"
)

type BodylinkEmail struct {
	URL string
}

func main() {

	templateData := BodylinkEmail{
		URL: "https://detik.id/",
	}
	to := "xxx@gmail.com"
	runtime.GOMAXPROCS(1)
	go SendEmailVerification(to, templateData)

	fmt.Println("kirim email sedang di proses, cek email anda")
}
func SendEmail(to string, subject string, data interface{}, templateFile string) error {
	result, _ := ParseTemplate(templateFile, data)
	m := gomail.NewMessage()
	m.SetHeader("From", "activation@xxx.com")
	m.SetHeader("To", to)
	// m.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", result)
	// m.Attach(templateFile) // attach whatever you want
	senderPort := 587
	d := gomail.NewDialer("mail.xxx.com", senderPort, "activation@xxx.com", "xxx")
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
	return err
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

func SendEmailVerification(to string, data interface{}) {
	var err error
	template := "email_template_verifikasi.html"
	subject := "sample email"
	err = SendEmail(to, subject, data, template)
	if err == nil {
		fmt.Println("send email '" + subject + "' success")
	} else {
		fmt.Println(err)
	}
}
