package providers

type DummyProvider struct {
	ID ProviderId
}

func NewDummyProvider() DummyProvider {
	return DummyProvider{ID: DummyProviderID}
}

func (provider DummyProvider) FetchOrders() error {
	return nil
}

func (provider DummyProvider) StopFetchOrders() error {
	return nil
}
