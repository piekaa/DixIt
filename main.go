package main

import (
	"dixit/web"
	"os"
)

func main() {

	port := "6999"
	externalHost := "192.168.1.105"
	externalPort := port

	args := os.Args[1:]

	if len(args) >= 1 {
		port = args[0]
		externalPort = port
	}

	if len(args) >= 2 {
		externalHost = args[1]
		externalPort = args[2]
	}

	web.NewServer(port, externalHost, externalPort)
}
