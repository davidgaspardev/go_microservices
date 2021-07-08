package service

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (service *Service) Run() {

	service.name = "Test"
	service.port = ":8080"

	mux := http.NewServeMux()

	// Handling with routes
	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "text/plain; charset=utf8")
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("Testing microservices with Golang"))

		fmt.Println("\n=> HTTP header (", time.Now().Local(), ")")
		fmt.Println("- Method:", request.Method)
		fmt.Println("- Path:", request.URL.Path)
		fmt.Println("- Query:", request.URL.RawQuery)
		fmt.Println("- ContentLength:", request.ContentLength)
		fmt.Println("- RemoteAddr:", request.RemoteAddr)
	})

	err := http.ListenAndServe(service.port, mux)
	if err != nil {
		log.Fatalf("Server dailed to start: %v", err)
	}
}

// func NewServer(mux *http.ServeMux) *http.Server {
// 	tlsConfig := &tls.Config{
// 		// See more here https://www.namecheap.com/blog/beginners-guide-to-tls-cipher-suites/
// 		PreferServerCipherSuites: false,
// 		CurvePreferences: []tls.CurveID{
// 			tls.CurveP256,
// 			tls.X25519, // Go 1.8 only
// 		},
// 		MinVersion: tls.VersionTLS12,
// 		CipherSuites: []uint16{
// 			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
// 			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
// 			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
// 			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
// 			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
// 			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
// 		},
// 	}

// 	server := &http.Server{
// 		Addr:         ":8080",
// 		ReadTimeout:  5 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 		IdleTimeout:  120 * time.Second,
// 		TLSConfig:    tlsConfig,
// 		Handler:      mux,
// 	}

// 	return server
// }
