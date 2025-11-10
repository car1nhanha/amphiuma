package utils

import (
	"encoding/base64"
	"log"
	"strings"
)

func DecodeBase64(file string) string {
	clean := strings.ReplaceAll(file, "\n", "")
	decoded, err := base64.StdEncoding.DecodeString(clean)
	if err != nil {
		log.Printf("erro ao decodificar base64: %v", err)
		return ""
	}
	return string(decoded)
}
