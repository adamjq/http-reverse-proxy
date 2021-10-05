//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/adamjq/http-reverse-proxy/internal/config"
	"github.com/adamjq/http-reverse-proxy/internal/proxy"
)

func CreateProxyServer() *proxy.ProxyServer {
	wire.Build(config.NewConfig, proxy.NewProxyServer)
	return &proxy.ProxyServer{}
}
