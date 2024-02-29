package server

import (
	"encoding/json"
	"github.com/a-h/templ"
	"log"
	"fmt"
	"net/http"
	"rabbitmq-blueprint/cmd/web"
	"rabbitmq-blueprint/internal/rabbitmq"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HelloWorldHandler)

	mux.HandleFunc("/health", s.healthHandler)
	mux.HandleFunc("/publish", s.publishHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/js/", fileServer)
	mux.Handle("/web", templ.Handler(web.HelloForm()))
	mux.HandleFunc("/hello", web.HelloWebHandler)

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) publishHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    message := r.FormValue("message")
    if message == "" {
        http.Error(w, "Message cannot be empty", http.StatusBadRequest)
        return
    }

    err := rabbitmq.SubmitMessage(message)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to publish message: %v", err), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Message '%s' published successfully", message)
}