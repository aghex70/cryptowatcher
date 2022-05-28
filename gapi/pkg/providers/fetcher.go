package providers

import (
	"gapi-agp/infrastructure/config"
)

type FetcherProvider struct {
	ID     ProviderId
	config config.FetcherConfig
}

func NewFetcherProvider() FetcherProvider {
	return FetcherProvider{ID: FetcherProviderID}
}

func (provider FetcherProvider) FetchOrders() error {
	return nil
}

func (provider FetcherProvider) StopFetchOrders() error {
	return nil
}
