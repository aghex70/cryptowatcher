package providers

type ProviderId int

const (
	DummyProviderID ProviderId = iota
	FetcherProviderID
)

type Provider interface {
	FetchOrders() error
	StopFetchOrders() error
}

func NewProvider(providerID ProviderId) (Provider, error) {
	switch providerID {
	case DummyProviderID:
		return NewDummyProvider(), nil
	case FetcherProviderID:
		return NewFetcherProvider(), nil
	default:
		return nil, ErrProviderNotFound
	}
}
