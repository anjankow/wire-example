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

var azureClientProviderSet wire.ProviderSet = wire.NewSet(
	azure.NewAzureClient,
	azure.NewAzureClientMock,

	wire.Bind(new(azure.AzureClientIfc), new(azure.AzureClient)),
	wire.Bind(new(azure.AzureClientIfc), new(azure.AzureClientMock)),
)

func ProvideAzure(cfg config.Config) azure.AzureService {
	panic(wire.Build(azureClientProviderSet, cfg))
}

func ProvideServer(cfg config.Config) *server.Server {
	panic(wire.Build(ProvideAzure, ProvideDB))
}

func ProvideCronJob(cfg config.Config) *server.Server {
	panic(wire.Build(ProvideAzure, ProvideDB))
}
