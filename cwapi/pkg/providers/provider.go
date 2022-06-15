package providers

import (
	"cwapi/config"
	"cwapi/pkg/errors"
	"go.uber.org/zap"
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
	logger          *zap.Logger
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
