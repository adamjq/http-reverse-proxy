// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/adamjq/http-reverse-proxy/internal/config"
	"github.com/adamjq/http-reverse-proxy/internal/proxy"
)

// Injectors from wire.go:

func CreateProxyServer() *proxy.ProxyServer {
	configConfig := config.NewConfig()
	proxyServer := proxy.NewProxyServer(configConfig)
	return proxyServer
}