package server

import (
	"database/sql"
	"wire-example/azure"
)

type Server struct {
	azure       azure.Azure
	azureClient azure.AzureClient
	db          *sql.DB
}

type CronJob struct {
	azure       azure.Azure
	azureClient azure.AzureClient
	db          *sql.DB
}
