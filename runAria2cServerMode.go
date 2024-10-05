package main

import (
	"log"
	"os"
	"os/exec"
)

func runAria2cServerMode() {
	// aria2c --enable-rpc --rpc-listen-all --rpc-secret=SECRET --listen-port=6900 --dht-listen-port=6900 --dir=/media/Downloads

	cmd := exec.Command("aria2c", "--enable-rpc", "--rpc-listen-all", "--rpc-secret="+aria2cRpcSecret, "--listen-port="+aria2cPort, "--dht-listen-port="+aria2cPort, "--dir="+aria2cDir)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	defer cmd.Process.Release()
}
