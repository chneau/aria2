package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func runAriaForever() {
	for {
		cmd := exec.Command("aria2c", "--enable-rpc", "--rpc-listen-all", "--rpc-secret="+secret, "--listen-port="+ariaPort, "--dht-listen-port="+ariaPort, "--dir="+ariaDir)
		cmd.Stdout = os.Stdout
		cmd.Run()
		log.Println("Restarting aria2c in 1s")
		<-time.After(1 * time.Second)
	}
}
