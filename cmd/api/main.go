package main

import (
	"rabbitmq-blueprint/internal/server"
	l "rabbitmq-blueprint/internal/logger"
)

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		l.Log.Fatalf("cannot start server: %s", err)
	}
}
