package emailProvider

import (
	"gopkg.in/gomail.v2"
	"os"
	"github.com/joho/godotenv"
	"strconv"
	"bytes"
	"html/template"

)


func SendEmail(subject string, to []string, filePath string, data map[string]interface{}) error {
	godotenv.Load()

	templateBytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	templateContent := string(templateBytes) // Convert []byte to string

	// Parse the HTML template
	tmpl, err := template.New("email").Parse(templateContent)
	if err != nil {
		return err
	}

	// Execute template with provided data
	var bodyBuffer bytes.Buffer
	if err := tmpl.Execute(&bodyBuffer, data); err != nil {
		return err
	}



	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return err
	}


	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", to...)
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", bodyBuffer.String())
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USER") , os.Getenv("SMTP_PASS"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil

}