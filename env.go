package main

import (
	_ "embed"
	"encoding/base64"
	"log"
	"os"
	"strings"
)

//go:embed public/index.html
var indexHtml string
var port = getEnvOrDefault("PORT", "3000")
var externalPort = getEnvOrDefault("EXTERNAL_PORT", port)
var defaultRpcSecret = randString(16)
var aria2cRpcSecret = getEnvOrDefault("ARIA2C_RPC_SECRET", defaultRpcSecret)
var aria2cPort = getEnvOrDefault("ARIA2C_PORT", "6800")
var aria2cDir = getEnvOrDefault("ARIA2C_DIR", "/home/c/data")

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
	encodedSecret := base64.StdEncoding.EncodeToString([]byte(aria2cRpcSecret))
	indexHtml = strings.Replace(indexHtml, `secret:""`, `secret:"`+encodedSecret+`"`, -1)
}
