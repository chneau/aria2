package main

import (
	_ "embed"
	"encoding/base64"
	"log"
	"os"
	"strings"
)

//go:embed aria-ng.html
var indexHtml string
var port = getEnvOrDefault("PORT", "3000")
var externalPort = getEnvOrDefault("EXTERNAL_PORT", port)
var defaultSecret = randString(16)
var secret = getEnvOrDefault("SECRET", defaultSecret)
var ariaPort = getEnvOrDefault("ARIA_PORT", "6800")
var ariaDir = getEnvOrDefault("ARIA_DIR", "/data")

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	log.Println(key+":", value)
	return value
}

func init() {
	indexHtml = strings.Replace(indexHtml, "6800", externalPort, -1)
	encodedSecret := base64.StdEncoding.EncodeToString([]byte(secret))
	indexHtml = strings.Replace(indexHtml, `secret:""`, `secret:"`+encodedSecret+`"`, -1)
}
