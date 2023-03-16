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
	fmt.Println("server azure client: ", server.Azure.GetClientName())
	if err := server.Azure.CheckDBConnection(); err != nil {
		panic(err)
	}

	serverMock, err := wiregen.ProvideServerMock(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("server mock azure client: ", serverMock.Azure.GetClientName())
	if err := serverMock.Azure.CheckDBConnection(); err != nil {
		panic(err)
	}

	cronJob, err := wiregen.ProvideCronJob(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("cronjob azure client: ", cronJob.Azure.GetClientName())
	if err := cronJob.Azure.CheckDBConnection(); err != nil {
		panic(err)
	}

	cronJobMock, err := wiregen.ProvideCronJob(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("cronjob mock azure client: ", cronJobMock.Azure.GetClientName())
	if err := cronJobMock.Azure.CheckDBConnection(); err != nil {
		panic(err)
	}

}
