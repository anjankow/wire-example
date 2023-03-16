// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wiregen

import (
	"github.com/google/wire"
	"wire-example/internal/azure"
	"wire-example/internal/config"
	"wire-example/internal/server"
)

// Injectors from wire.go:

func ProvideAzureClient(cfg config.Config) azure.AzureClientIfc {
	azureClient := azure.NewAzureClient(cfg)
	return azureClient
}

func ProvideAzureClientMock(cfg config.Config) azure.AzureClientIfc {
	azureClientMock := azure.NewAzureClientMock(cfg)
	return azureClientMock
}

func ProvideServer(cfg config.Config) (*server.Server, error) {
	azureClientIfc := ProvideAzureClient(cfg)
	db, err := server.NewDB(cfg)
	if err != nil {
		return nil, err
	}
	azureService, err := azure.NewAzureService(cfg, azureClientIfc, db)
	if err != nil {
		return nil, err
	}
	serverServer, err := server.NewServer(cfg, azureService, azureClientIfc, db)
	if err != nil {
		return nil, err
	}
	return serverServer, nil
}

func ProvideServerMock(cfg config.Config) (*server.Server, error) {
	azureClientIfc := ProvideAzureClientMock(cfg)
	db, err := server.NewDB(cfg)
	if err != nil {
		return nil, err
	}
	azureService, err := azure.NewAzureService(cfg, azureClientIfc, db)
	if err != nil {
		return nil, err
	}
	serverServer, err := server.NewServer(cfg, azureService, azureClientIfc, db)
	if err != nil {
		return nil, err
	}
	return serverServer, nil
}

func ProvideCronJob(cfg config.Config) (*server.Server, error) {
	azureClientIfc := ProvideAzureClient(cfg)
	db, err := server.NewDB(cfg)
	if err != nil {
		return nil, err
	}
	azureService, err := azure.NewAzureService(cfg, azureClientIfc, db)
	if err != nil {
		return nil, err
	}
	serverServer, err := server.NewServer(cfg, azureService, azureClientIfc, db)
	if err != nil {
		return nil, err
	}
	return serverServer, nil
}

func ProvideCronJobMock(cfg config.Config) (*server.CronJob, error) {
	azureClientIfc := ProvideAzureClientMock(cfg)
	db, err := server.NewDB(cfg)
	if err != nil {
		return nil, err
	}
	azureService, err := azure.NewAzureService(cfg, azureClientIfc, db)
	if err != nil {
		return nil, err
	}
	cronJob, err := server.NewCronJob(cfg, azureService, azureClientIfc, db)
	if err != nil {
		return nil, err
	}
	return cronJob, nil
}

// wire.go:

var azureClientSet wire.ProviderSet = wire.NewSet(azure.NewAzureClient, wire.Bind(new(azure.AzureClientIfc), new(azure.AzureClient)))

var azureClientMockSet wire.ProviderSet = wire.NewSet(azure.NewAzureClientMock, wire.Bind(new(azure.AzureClientIfc), new(azure.AzureClientMock)))
