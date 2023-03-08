package server

import (
	"database/sql"
	"wire-example/internal/azure"
	"wire-example/internal/config"

	"github.com/DATA-DOG/go-sqlmock"
)

type Server struct {
	Azure       *azure.AzureService
	azureClient azure.AzureClientIfc
	db          *sql.DB
	cfg         config.Config
}

type CronJob struct {
	Azure       *azure.AzureService
	azureClient azure.AzureClientIfc
	db          *sql.DB
	cfg         config.Config
}

func NewDB(cfg config.Config) (*sql.DB, error) {
	db, _, err := sqlmock.New()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewServer(cfg config.Config, azureService *azure.AzureService, azureClient azure.AzureClientIfc, db *sql.DB) (*Server, error) {
	return &Server{
		db:          db,
		azureClient: azureClient,
		cfg:         cfg,
		Azure:       azureService,
	}, nil
}
