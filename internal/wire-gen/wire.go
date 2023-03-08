//go:build wireinject
// +build wireinject

package wiregen

import (
	"wire-example/internal/azure"
	"wire-example/internal/config"
	"wire-example/internal/server"

	"github.com/google/wire"
)

var azureClientSet wire.ProviderSet = wire.NewSet(
	azure.NewAzureClient,
	wire.Bind(new(azure.AzureClientIfc), new(azure.AzureClient)),
)
var azureClientMockSet wire.ProviderSet = wire.NewSet(
	azure.NewAzureClientMock,
	wire.Bind(new(azure.AzureClientIfc), new(azure.AzureClientMock)),
)

func ProvideAzureClient(cfg config.Config) azure.AzureClientIfc {
	panic(wire.Build(azureClientSet))
	// return new(azure.AzureClient), nil
}

func ProvideAzureClientMock(cfg config.Config) azure.AzureClientIfc {
	panic(wire.Build(azureClientMockSet))
	// return new(azure.AzureClientMock), nil
}

func ProvideServer(cfg config.Config) (*server.Server, error) {
	wire.Build(server.NewServer, ProvideAzureClient, azure.NewAzureService, server.NewDB)
	return new(server.Server), nil
}

func ProvideServerMock(cfg config.Config) (*server.Server, error) {
	wire.Build(server.NewServer, ProvideAzureClientMock, azure.NewAzureService, server.NewDB)
	return new(server.Server), nil
}
