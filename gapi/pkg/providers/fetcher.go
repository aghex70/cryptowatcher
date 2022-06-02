package providers

import (
	"gapi-agp/config"
	"go.uber.org/zap"
)

type FetcherProvider struct {
	ID     ProviderId
	config config.FetcherConfig
	logger *zap.Logger
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
