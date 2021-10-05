package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/adamjq/http-reverse-proxy/internal/config"
	"github.com/go-chi/chi"
)

type ProxyServer struct {
	cfg *config.Config
}

func NewProxyServer(cfg *config.Config) *ProxyServer {
	return &ProxyServer{
		cfg: cfg,
	}
}

func (ps *ProxyServer) Start() error {
	log.Printf("Proxy is listening on :%s\n", ps.cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", ps.cfg.Port), ps.proxyHandler())
}

func (ps *ProxyServer) proxyHandler() http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Handle("/service_a/*", proxyrequest(ps.cfg.ServiceAUrl))
		r.Handle("/service_b/*", proxyrequest(ps.cfg.ServiceBUrl))
	})

	return r
}

func proxyrequest(uri string) http.Handler {
	u, err := url.ParseRequestURI(uri)
	if err != nil {
		panic(err)
	}
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = u.Scheme
			req.URL.Host = u.Host
		},
	}
}
