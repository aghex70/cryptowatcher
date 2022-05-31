package providers

import (
	"gapi-agp/infrastructure/config"
	"gapi-agp/pkg/errors"
)

type ProviderId int

const (
	DummyProviderID ProviderId = iota
	FetcherProviderID
)

type Provider interface {
	FetchOrders() error
	StopFetchOrders() error
}

type ProviderManager struct {
	providersConfig config.ProvidersConfig
}

func NewProviderManager() *ProviderManager {
	return &ProviderManager{providersConfig: config.C.Providers}
}

func (pm ProviderManager) GetProvider(pid ProviderId) (Provider, error) {
	switch pid {
	case DummyProviderID:
		return NewDummyProvider(), nil
	case FetcherProviderID:
		return NewFetcherProvider(), nil
	default:
		return nil, errors.ErrProviderNotFound
	}
}
