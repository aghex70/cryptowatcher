package ports

import "gapi-agp/internal/core/domain"

type TradeRepository interface {
	GetTrades() ([]domain.Trade, error)
	GetTradesBySymbol(symbol string) ([]domain.Trade, error)
	GetTradesByEventType(eventType domain.EventType) ([]domain.Trade, error)
	GetTradesByUserIdAndExternalId(map[string]int) ([]domain.Trade, error)
}

type UserRepository interface {
	GetUser(userID int) (domain.User, error)
}

type CacheRepository interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
}
