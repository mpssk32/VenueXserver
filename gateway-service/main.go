package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func newProxy(target string) *httputil.ReverseProxy {
	url, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	return proxy
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// public routes
		if strings.HasPrefix(r.URL.Path, "/auth/login") ||
			strings.HasPrefix(r.URL.Path, "/auth/register") {

			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(
				w,
				"missing authorization header",
				http.StatusUnauthorized,
			)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(
				w,
				"invalid token",
				http.StatusUnauthorized,
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	// PROXIES
	authProxy := newProxy(
		"http://auth-service:8080",
	)

	concertProxy := newProxy(
		"http://concert-service:8082",
	)

	chatProxy := newProxy(
		"http://chat-service:8081",
	)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		log.Printf(
			"[%s] %s %s",
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
		)

		// HEALTHCHECK
		if r.URL.Path == "/health" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
			return
		}

		// CORS
		w.Header().Set(
			"Access-Control-Allow-Origin",
			"*",
		)

		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization",
		)

		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, DELETE, OPTIONS",
		)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// AUTH SERVICE
		if strings.HasPrefix(r.URL.Path, "/auth") {

			authProxy.ServeHTTP(w, r)
			return
		}

		// CHAT SERVICE
		if strings.HasPrefix(r.URL.Path, "/chat") {

			r.URL.Path = strings.TrimPrefix(
				r.URL.Path,
				"/chat",
			)

			authMiddleware(chatProxy).ServeHTTP(w, r)
			return
		}

		// CONCERT SERVICE
		if strings.HasPrefix(r.URL.Path, "/events") ||
			strings.HasPrefix(r.URL.Path, "/concerts") ||
			strings.HasPrefix(r.URL.Path, "/tickets") ||
			strings.HasPrefix(r.URL.Path, "/applications") ||
			strings.HasPrefix(r.URL.Path, "/venues") {

			authMiddleware(concertProxy).ServeHTTP(w, r)
			return
		}

		// NOT FOUND
		http.NotFound(w, r)
	})

	server := &http.Server{
		Addr:         ":8000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Gateway started on :8000")

	log.Fatal(
		server.ListenAndServe(),
	)
}