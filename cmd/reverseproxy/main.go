package main

import (
	"log"

	"github.com/adamjq/http-reverse-proxy/internal/config"
	"github.com/adamjq/http-reverse-proxy/internal/proxy"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	proxyserver := proxy.NewProxyServer(cfg)

	err = proxyserver.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Exiting...\n")
}
