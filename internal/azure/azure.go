package azure

import (
	"database/sql"
	"wire-example/internal/config"
)

type AzureService struct {
	client AzureClientIfc
	db     *sql.DB
}

func NewAzureService(cfg config.Config, client AzureClientIfc, db *sql.DB) (*AzureService, error) {
	return &AzureService{
		client: client,
		db:     db,
	}, nil
}

func (a AzureService) GetClientName() string {
	return a.client.GetName()
}

func (a AzureService) CheckDBConnection() error {
	return a.db.Ping()
}
