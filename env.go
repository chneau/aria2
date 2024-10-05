package main

import (
	"log"
	"os"
)

var port = getEnvOrDefault("PORT", "3000")
var externalPort = getEnvOrDefault("EXTERNAL_PORT", port)
var defaultRpcSecret = randString(16)
var aria2cRpcSecret = getEnvOrDefault("ARIA2C_RPC_SECRET", defaultRpcSecret)
var aria2cPort = getEnvOrDefault("ARIA2C_PORT", "6800")
var aria2cDir = getEnvOrDefault("ARIA2C_DIR", "/media/Downloads")

func init() {
	log.Println("PORT:", port)
	log.Println("EXTERNAL_PORT:", externalPort)
	log.Println("ARIA2C_RPC_SECRET:", aria2cRpcSecret)
	log.Println("ARIA2C_PORT:", aria2cPort)
	log.Println("ARIA2C_DIR:", aria2cDir)
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}
