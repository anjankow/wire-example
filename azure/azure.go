package azure

import "database/sql"

type AzureService struct {
	client AzureClientIfc
	db     *sql.DB
}

func (a AzureService) GetClientName() string {
	return a.client.GetName()
}

func (a AzureService) CheckDBConnection() error {
	return a.db.Ping()
}
