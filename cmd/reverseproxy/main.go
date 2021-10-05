package main

import (
	"log"
)

func main() {
	proxyserver := CreateProxyServer()

	err := proxyserver.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Exiting...\n")
}
