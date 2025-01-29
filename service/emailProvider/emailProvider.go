package emailProvider

import (
	"gopkg.in/gomail.v2"
)

func SendEmail(subject string, to []string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", to...)
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("sandbox.smtp.mailtrap.io", 2525, "ba711fbbe0c962", "561740d0aaff95")

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil

}