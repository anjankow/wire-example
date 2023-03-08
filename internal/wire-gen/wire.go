//+build wireinject

package wiregen

import (
	"wire-example/internal/azure"
	"wire-example/internal/config"
	"wire-example/internal/server"

	"github.com/google/wire"
)

var azureSet wire.ProviderSet = wire.NewSet(
	azure.NewAzureClient,
	wire.Bind(new(azure.AzureClientIfc), new(azure.AzureClient)),
)
var azureMockSet wire.ProviderSet = wire.NewSet(
	azure.NewAzureClientMock,
	wire.Bind(new(azure.AzureClientIfc), new(azure.AzureClientMock)),
)

func ProvideAzure() (*azure.AzureService, error) {
	wire.Build(azureSet)
	return new(azure.AzureService), nil
}

func ProvideAzureMock() (*azure.AzureService, error) {
	wire.Build(azureMockSet)
	return new(azure.AzureService), nil
}

func ProvideServer(cfg config.Config) (*server.Server, error) {
	if cfg.UseAzureMock {
		wire.Build(ProvideAzureMock, server.NewDB)
	} else {
		wire.Build(ProvideAzureMock, server.NewDB)
	}

	return new(server.Server), nil
}

func ProvideCronJob(cfg config.Config) (*server.Server, error) {
	if cfg.UseAzureMock {
		wire.Build(ProvideAzureMock, server.NewDB)
	} else {
		wire.Build(ProvideAzureMock, server.NewDB)
	}

	return new(server.Server), nil
}
