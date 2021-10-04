package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	handler := proxyhandler()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Printf("Proxy is listening on :%s\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Proxy is exiting...\n")
}

func proxyhandler() http.Handler {
	r := chi.NewRouter()

	serviceA := os.Getenv("SERVICE_A_URL")
	if serviceA == "" {
		serviceA = "http://localhost:5001"
	}

	serviceB := os.Getenv("SERVICE_B_URL")
	if serviceB == "" {
		serviceB = "http://localhost:5002"
	}

	r.Group(func(r chi.Router) {
		r.Handle("/service_a/*", proxyrequest(serviceA))
		r.Handle("/service_b/*", proxyrequest(serviceB))
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
