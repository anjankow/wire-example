package mywire

import (
	"database/sql"
	"wire-example/azure"
	"wire-example/config"
	"wire-example/server"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/wire"
)

func ProvideDB(cfg config.Config) *sql.DB {
	db, _, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	return db
}

var azureSet wire.ProviderSet = wire.NewSet(
	azure.NewAzureClient,
	wire.Bind(new(azure.AzureClientIfc), new(azure.AzureClient)),
)
var azureMockSet wire.ProviderSet = wire.NewSet(
	azure.NewAzureClientMock,
	wire.Bind(new(azure.AzureClientIfc), new(azure.AzureClientMock)),
)

func ProvideAzure() azure.AzureService {
	panic(wire.Build(azureSet))
}

func ProvideAzureMock() azure.AzureService {
	panic(wire.Build(azureMockSet))
}

func ProvideServer(cfg config.Config) *server.Server {
	if cfg.UseAzureMock {
		panic(wire.Build(ProvideAzureMock, ProvideDB))
	} else {
		panic(wire.Build(ProvideAzureMock, ProvideDB))
	}
}

func ProvideCronJob(cfg config.Config) *server.Server {
	if cfg.UseAzureMock {
		panic(wire.Build(ProvideAzureMock, ProvideDB))
	} else {
		panic(wire.Build(ProvideAzureMock, ProvideDB))
	}
}
