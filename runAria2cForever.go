package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func runAria2cForever() {
	for {
		cmd := exec.Command("aria2c", "--enable-rpc", "--rpc-listen-all", "--rpc-secret="+aria2cRpcSecret, "--listen-port="+aria2cPort, "--dht-listen-port="+aria2cPort, "--dir="+aria2cDir)
		cmd.Stdout = os.Stdout
		cmd.Run()
		log.Println("Restarting aria2c in 1s")
		<-time.After(1 * time.Second)
	}
}
