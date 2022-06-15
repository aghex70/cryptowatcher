package ports

import (
	"cwapi/internal/core/domain"
	"cwapi/pkg/providers"
)

type TradeUseCase interface {
	Fetch(r FetchRequest) (*providers.FetchTradesResponse, error)
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
