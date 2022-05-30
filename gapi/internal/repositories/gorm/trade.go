package gorm

import (
	"gapi-agp/internal/core/domain"
	"time"
)

type Trade struct {
	ID         int              `gorm:"coulmn:id;type:int;auto_increment;primary_key"`
	EventType  domain.EventType `gorm:"column:event_type"`
	ExternalID int              `gorm:"column:external_id"`
	Symbol     string           `gorm:"column:symbol"`
	Price      float32          `gorm:"column:price"`
	Currency   string           `gorm:"column:currency"`
	Quantity   float32          `gorm:"column:quantity"`
	EventTime  time.Time        `gorm:"column:event_time"`
	TradeTime  time.Time        `gorm:"column:trade_time"`
	BuyerID    int              `gorm:"column:buyer_id"`
	SellerID   int              `gorm:"column:seller_id"`
	Source     string           `gorm:"column:source"`
	CreatedAt  time.Time        `gorm:"column:created_at"`
	UpdatedAt  time.Time        `gorm:"column:updated_at"`
}

func (g GormRepo) GetTrades() ([]domain.Trade, error) {
	var trades []domain.Trade
	tx := g.DB.Order("event_time").Find(&trades)
	if tx.Error != nil {
		return []domain.Trade{}, tx.Error
	}
	return trades, nil
}

func (g GormRepo) GetTradesBySymbol(symbol string) ([]domain.Trade, error) {
	var trades []domain.Trade
	tx := g.DB.Where(&Trade{Symbol: symbol}).Find(&trades)
	if tx.Error != nil {
		return []domain.Trade{}, tx.Error
	}
	return trades, nil
}

func (g GormRepo) GetTradesByEventType(eventType domain.EventType) ([]domain.Trade, error) {
	var trades []domain.Trade
	tx := g.DB.Where("event_type = ?", eventType).Find(&trades)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return trades, nil
}

func (g GormRepo) GetTradesByUserIdAndExternalId(fields map[string]int) ([]domain.Trade, error) {
	var trades []domain.Trade
	tx := g.DB.Where(fields).Find(&trades)
	if tx.Error != nil {
		return []domain.Trade{}, tx.Error
	}
	return trades, nil
}
