package main

import "os"

var port = getEnvOrDefault("PORT", "3000")
var externalPort = getEnvOrDefault("EXTERNAL_PORT", port)
var defaultRpcSecret = randString(16)
var rpcSecret = getEnvOrDefault("RPC_SECRET", defaultRpcSecret)

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}
