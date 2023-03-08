package azure

import "wire-example/internal/config"

type AzureClientIfc interface {
	GetName() string
}

type AzureClient struct {
	name        string
	ConfigValue string
}

type AzureClientMock struct {
	mockName    string
	ConfigValue string
}

func NewAzureClient(cfg config.Config) AzureClient {
	return AzureClient{
		name:        "client",
		ConfigValue: cfg.ImportantConfigValue}
}

func NewAzureClientMock(cfg config.Config) AzureClientMock {
	return AzureClientMock{
		mockName:    "mock",
		ConfigValue: cfg.ImportantConfigValue}
}

func (a AzureClient) GetName() string {
	return a.name
}

func (a AzureClientMock) GetName() string {
	return a.mockName
}
