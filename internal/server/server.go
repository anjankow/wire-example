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
	Cfg         config.Config
}

type CronJob struct {
	Azure       *azure.AzureService
	azureClient azure.AzureClientIfc
	db          *sql.DB
	Cfg         config.Config
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
		Cfg:         cfg,
		Azure:       azureService,
	}, nil
}

func NewCronJob(cfg config.Config, azureService *azure.AzureService, azureClient azure.AzureClientIfc, db *sql.DB) (*CronJob, error) {
	return &CronJob{
		db:          db,
		azureClient: azureClient,
		Cfg:         cfg,
		Azure:       azureService,
	}, nil
}
