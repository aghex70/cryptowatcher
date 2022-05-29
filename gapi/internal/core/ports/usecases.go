package ports

import "gapi-agp/internal/core/domain"

type TradeUseCase interface {
	Fetch(r FetchRequest) error
	StopFetch(r StopFetchRequest) error
	Get() ([]domain.Trade, error)
	GetSales() ([]domain.Trade, error)
	GetPurchases() ([]domain.Trade, error)
}

type UserUseCase interface {
	Get(userID int) (domain.User, error)
	GetUserTrades(userID int) ([]domain.Trade, error)
}

type FetchRequest struct {
	Source string
}

type StopFetchRequest struct {
	Source string
}
