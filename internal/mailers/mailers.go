package mailers

import (
	"asyncfiber/internal/config"
	"bytes"
	"log"
	"text/template"

	"gopkg.in/gomail.v2"
)

func SendHTML(to string) error {
	t, err := template.ParseFiles("pkg/template/email_test.html")
	if err != nil {
		log.Fatal(err)
	}
	var body bytes.Buffer
	t.Execute(&body, struct{ Name string }{Name: "Kawasaki"})

	mail := gomail.NewMessage()
	mail.SetHeader("From", config.Email)
	mail.SetHeader("To", to)
	// mail.SetAddressHeader("Cc", "dan@example.com", "Dan")
	mail.SetHeader("Subject", "Thư xác nhận sử dụng phòng Lab")
	mail.SetBody("text/html", body.String())
	// mail.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, config.Email, config.EmailAppPassword)
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}
