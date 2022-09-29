package utils

import (
	"os"
	"strings"
)

func ParsePort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	return port
}
