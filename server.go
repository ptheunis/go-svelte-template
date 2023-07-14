package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ptheunis/go-svelte/ui"
)

const (
	writeTimeout = (15 * time.Second)
	readTimeout  = (15 * time.Second)
	idleTimeout  = (60 * time.Second)
)

var hostAddr = ":8080"

// GetServer ...
func GetServer(listenPort int) (*http.Server, error) {
	assets, _ := ui.Assets()

	// Use the file system to serve static files
	fs := http.FileServer(http.FS(assets))
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(fs)

	/*
		// generate certificate pair
		cert, err := tls.LoadX509KeyPair("keys/localhost.crt", "keys/localhost.key")
		if err != nil {
			return nil, err
		}

		// create the TLS Config with the CA pool and enable Client certificate validation
		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
			MinVersion:   tls.VersionTLS13,
		}
	*/

	// Config of session/cookie settings
	// s := sessions.GetInstance()
	// s.SetDomain("<realdomain>")
	// s.SetExpirationDuration(1 * time.Hour)

	return &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", listenPort),
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
		//TLSConfig:    tlsConfig,
	}, nil
}
