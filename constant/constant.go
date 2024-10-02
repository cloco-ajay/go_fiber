package constant

import "fmt"

const (
	PORT = "7000"
	HOST = "http://127.0.0.1"
)

func GetBaseURL() string {
	return fmt.Sprintf("%s:%s", HOST, PORT)
}
