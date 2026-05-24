package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {

	authURL, _ := url.Parse(
		"http://auth-service:8080",
	)

	chatURL, _ := url.Parse(
		"http://chat-service:8081",
	)

	authProxy := httputil.NewSingleHostReverseProxy(
		authURL,
	)

	chatProxy := httputil.NewSingleHostReverseProxy(
		chatURL,
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		log.Printf(
			"%s %s",
			r.Method,
			r.URL.Path,
		)

		// CHAT ROUTES
		if strings.HasPrefix(r.URL.Path, "/chat") {

			r.URL.Path = strings.TrimPrefix(
				r.URL.Path,
				"/chat",
			)

			chatProxy.ServeHTTP(w, r)

			return
		}

		// AUTH ROUTES
		authProxy.ServeHTTP(w, r)
	})

	log.Println("Gateway started on :8000")

	log.Fatal(
		http.ListenAndServe(":8000", nil),
	)
}