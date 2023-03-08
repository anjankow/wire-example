package main

import (
	"fmt"
	"wire-example/internal/server"
)

func main() {
	// cfg := config.Config{UseAzureMock: true}

	server := server.Server{}
	fmt.Print(server.Azure.GetClientName())
	if err := server.Azure.CheckDBConnection(); err != nil {
		panic(err)
	}

}
