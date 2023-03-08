package main

import (
	"fmt"
	"wire-example/internal/config"
	wiregen "wire-example/internal/wire-gen"
)

func main() {
	cfg := config.Config{}

	server, err := wiregen.ProvideServer(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(server.Azure.GetClientName())
	if err := server.Azure.CheckDBConnection(); err != nil {
		panic(err)
	}

}
