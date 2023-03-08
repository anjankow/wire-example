package server

import (
	"database/sql"
	"wire-example/internal/azure"
	"wire-example/internal/config"

	"github.com/DATA-DOG/go-sqlmock"
)

type Server struct {
	Azure       azure.AzureService
	azureClient azure.AzureClient
	db          *sql.DB
}

type CronJob struct {
	Azure       azure.AzureService
	azureClient azure.AzureClient
	db          *sql.DB
}

func NewDB(cfg config.Config) (*sql.DB, error) {
	db, _, err := sqlmock.New()
	if err != nil {
		return nil, err
	}
	return db, nil
}
