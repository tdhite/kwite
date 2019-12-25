/*
server.go

Copyright (c) 2019-2020 VMware, Inc.

SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
*/
package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"time"
)

// We add some methods to http.Server, which requires adding a new type
type HttpServer struct {
	http.Server
}

// Try graceful shutdown (if possible)
func (s *HttpServer) shutdown(shutdownTimeout int, ossig <-chan os.Signal, finished chan<- struct{}) {
	// wait on signal to start the server shutdown
	<-ossig
	s.ErrorLog.Println("Recieved signal to stop, shutting down.")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(shutdownTimeout))
	defer cancel()

	s.SetKeepAlivesEnabled(false)
	if err := s.Shutdown(ctx); err != nil {
		s.ErrorLog.Fatalf("Gracefull shutdown failed: %s\n", err)
	}

	// write (via closing) the finished channel to free up anyone waiting
	close(finished)
}

// Start the server, i.e., accept requests and handle responses
// The shutdownTimeout provides a max time graceful shutdown is allowed to take
// before exiting the shutdown attempt regardless of success.
func (s *HttpServer) StartServer(shutdownTimeout int) error {
	ossig := make(chan os.Signal, 1)
	finished := make(chan struct{}, 1)

	signal.Notify(ossig, os.Interrupt)
	go s.shutdown(shutdownTimeout, ossig, finished)
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.ErrorLog.Println("Failed to start server ", err)
	}

	// wait 'til shutdown is finished
	<-finished
	s.ErrorLog.Println("Server succesfully shut down.")

	return err
}

// Allocate, initialize and return a new HttpServer
func New(tmpl Template, m map[string]string, addr string, l *log.Logger) *HttpServer {
	r := http.NewServeMux()
	url := m["url"]

	l.Println("Setting up server urls at: " + url)
	r.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		tmpl.HttpHandler(w, r)
	})
	r.HandleFunc(path.Join(url, "kwiteready"), func(w http.ResponseWriter, r *http.Request) {
		tmpl.HttpHandler(w, r)
	})
	r.HandleFunc(path.Join(url, "kwitealive"), func(w http.ResponseWriter, r *http.Request) {
		tmpl.HttpHandler(w, r)
	})
	// Add a catch-all to log the issues.
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.HttpUnknownHandler(w, r)
	})

	server := HttpServer{}
	server.Addr = addr
	server.Handler = r
	server.ErrorLog = l
	server.ReadTimeout = 5 * time.Second
	server.WriteTimeout = 10 * time.Second
	server.IdleTimeout = 15 * time.Second
	server.MaxHeaderBytes = 1 << 12

	return &server
}
