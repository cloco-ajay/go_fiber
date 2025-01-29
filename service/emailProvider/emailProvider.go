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
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return err
	}

	body,err := ExecuteHtemplate(filePath, data)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USER") , os.Getenv("SMTP_PASS"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil

}

func ExecuteHtemplate(filePath string, data map[string]interface{}) (string, error) {
	templateBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	templateContent := string(templateBytes)

	tmpl, err := template.New("email").Parse(templateContent)
	if err != nil {
		return "", err
	}

	var bodyBuffer bytes.Buffer
	if err := tmpl.Execute(&bodyBuffer, data); err != nil {
		return "", err
	}

	
	return bodyBuffer.String(), nil
}