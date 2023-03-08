package main

import (
	"fmt"
	"wire-example/config"
	"wire-example/mywire"
)

func main() {
	cfg := config.Config{UseAzureMock: true}

	server := mywire.ProvideServer(cfg)
	fmt.Print(server.Azure.GetClientName())
	if err := server.Azure.CheckDBConnection(); err != nil {
		panic(err)
	}

}
