package main

import (
	"net/http"
	"wire-example/config"
)

func main() {
	cfg := config.Config{UserAzureMock: true}

	server := server.Wire(cfg)

	userHandler := user.Wire(db)
	http.Handle("/user", userHandler.FetchByUsername())
	http.ListenAndServe(":8000", nil)
}
