package main

import (
	"log"
	"os"
	"os/exec"
)

func runAria2cServerMode() {
	cmd := exec.Command("aria2c", "--enable-rpc", "--rpc-listen-all", "--rpc-secret="+aria2cRpcSecret, "--listen-port="+aria2cPort, "--dht-listen-port="+aria2cPort, "--dir="+aria2cDir)
	cmd.Stdout = os.Stdout

	for {
		err := cmd.Start()
		if err == nil {
			break
		}
		log.Println(err)
	}

	defer cmd.Process.Release()
}
