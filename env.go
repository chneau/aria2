package main

import (
	_ "embed"
	"encoding/base64"
	"log"
	"os"
	"strings"
)

var (
	//go:embed aria-ng.html
	indexHtml    string
	port         = getEnvOrDefault("PORT", "3000")
	externalPort = getEnvOrDefault("EXTERNAL_PORT", port)
	secret       = getEnvOrDefault("SECRET", randString(16))
	ariaPort     = getEnvOrDefault("ARIA_PORT", "6800")
	ariaDir      = getEnvOrDefault("ARIA_DIR", "/data")
	username     = getEnvOrDefault("USERNAME", "")
	password     = getEnvOrDefault("PASSWORD", secret)
	authSet      = username != "" || password != ""
)

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
	if authSet {
		log.Println("Setting rpc secret on index.html since username or password are set")
		encodedSecret := base64.StdEncoding.EncodeToString([]byte(secret))
		indexHtml = strings.Replace(indexHtml, `secret:""`, `secret:"`+encodedSecret+`"`, -1)
	} else {
		log.Println("Not setting rpc secret on index.html since username and password are not set")
		log.Println("Confugure the rpc secret in the aria2 settings:", secret)
	}
}
