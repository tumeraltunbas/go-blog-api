package utils

import (
	"os"
)

func GetPort() string {
	port := os.Getenv("PORT")
	return ":" + port
}
