package server

import (
	"database/sql"
	"wire-example/azure"
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
