package webserver

import (
	"context"
	"fmt"
	"github.com/dzheremi/mqttviz/spotify"
	"html/template"
	"net/http"
)

// Start - start web server for authentication
func Start() {
	httpMux := http.NewServeMux()
	httpServer := http.Server{Addr: ":5555", Handler: httpMux}
	ctx, cancel := context.WithCancel(context.Background())
	httpMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template, err := template.ParseFiles("webserver/templates/landing.html")
		if err != nil {
			panic("--- Template not found ---")
		}
		template.Execute(w, spotify.ClientCredentials)
	})
	httpMux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["code"]
		if !ok || len(keys[0]) < 1 {
			panic("--- Invalid response from Spotify ---")
		}
		spotify.ClientCredentials.Code = keys[0]
		fmt.Fprintf(w, "Authentication Complete")
		cancel()
	})
	go func() {
		httpServer.ListenAndServe()
	}()
	<-ctx.Done()
	httpServer.Shutdown(ctx)
}

// API - start API server
func API() {
	httpMux := http.NewServeMux()
	httpServer := http.Server{Addr: ":5555", Handler: httpMux}
	httpMux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "{\"paused:\": %t}", spotify.Pause)
		}
	})
	httpMux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			spotify.Pause = true
			fmt.Fprintf(w, "{\"paused:\": %t}", spotify.Pause)
		}
	})
	httpMux.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			spotify.Pause = false
			fmt.Fprintf(w, "{\"paused:\": %t}", spotify.Pause)
		}
	})
	httpServer.ListenAndServe()
}
