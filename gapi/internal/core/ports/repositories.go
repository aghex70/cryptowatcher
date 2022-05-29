package ports

import "gapi-agp/internal/core/domain"

type TradeRepository interface {
	Get() ([]domain.Trade, error)
	GetBySymbol(symbol string) ([]domain.Trade, error)
	GetByEventType(eventType domain.EventType) ([]domain.Trade, error)
	GetByUserIdAndExternalId(map[string]int) ([]domain.Trade, error)
}

type UserRepository interface {
	Get(userID int) (domain.User, error)
}

type CacheRepository interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
}
