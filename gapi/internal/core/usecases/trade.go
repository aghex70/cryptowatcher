package usecases

import (
	"gapi-agp/internal/core/domain"
	"gapi-agp/internal/core/ports"
)

type TradeInteractor struct {
	tradeRepo ports.TradeRepository
}

func (interactor TradeInteractor) Fetch(source string) error {
	return nil
}

func (interactor TradeInteractor) StopFetch(source string) error{
	return nil
}

func (interactor TradeInteractor) Get() ([]domain.Trade, error){
	return []domain.Trade, nil
}

func (interactor TradeInteractor) GetSales() ([]domain.Trade, error){
	return []domain.Trade, nil
}

func (interactor TradeInteractor) GetPurchases() ([]domain.Trade, error){
	return []domain.Trade, nil
}

type FetchOrdersRequest struct {
	Source string
}

type StopFetchOrdersRequest struct {
	Source string
}