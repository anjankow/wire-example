package azure

import "database/sql"

type AzureService struct {
	client AzureClientIfc
	db     *sql.DB
}

func (a Azure) GetClientName() string {
	return a.client.GetName()
}

func (a Azure) CheckDBConnection() error {
	return a.db.Ping()
}
