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
}
