package gorm

import (
	"cwapi/internal/core/domain"
	"database/sql"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type TradeGormRepo struct {
	*gorm.DB
	SqlDB  *sql.DB
	logger *zap.Logger
}

type Trade struct {
	ID          int              `gorm:"coulmn:id;type:int;auto_increment;primary_key"`
	ProviderID  int              `gorm:"column:provider_id"`
	TradeID     int              `gorm:"column:trade_id"`
	SymbolsFrom string           `gorm:"column:symbols_from"`
	SymbolsTo   string           `gorm:"column:symbols_to"`
	Currency    string           `gorm:"column:currency"`
	Price       float32          `gorm:"column:price"`
	Quantity    float32          `gorm:"column:quantity"`
	TradeType   domain.TradeType `gorm:"column:trade_type"`
	EventTime   time.Time        `gorm:"column:event_time"`
	TradeTime   time.Time        `gorm:"column:trade_time"`
	BuyerID     int              `gorm:"column:buyer_id"`
	SellerID    int              `gorm:"column:seller_id"`
	CreatedAt   time.Time        `gorm:"column:created_at"`
	UpdatedAt   time.Time        `gorm:"column:updated_at"`
}

func NewTradeGormRepo(db *gorm.DB) (*TradeGormRepo, error) {
	return &TradeGormRepo{
		DB: db,
	}, nil
}

func (g TradeGormRepo) GetTrades() ([]domain.Trade, error) {
	var trades []domain.Trade
	tx := g.DB.Order("created_at").Find(&trades)
	if tx.Error != nil {
		return []domain.Trade{}, tx.Error
	}
	return trades, nil
}

func (g TradeGormRepo) GetTradesBySymbol(symbol string) ([]domain.Trade, error) {
	var trades []domain.Trade
	tx := g.DB.Where(&Trade{SymbolsFrom: symbol}).Or(&Trade{SymbolsTo: symbol}).Find(&trades)
	if tx.Error != nil {
		return []domain.Trade{}, tx.Error
	}
	return trades, nil
}

func (g TradeGormRepo) GetTradesByTradeType(TradeType domain.TradeType) ([]domain.Trade, error) {
	var trades []domain.Trade
	tx := g.DB.Where("trade_type = ?", TradeType).Find(&trades)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return trades, nil
}

func (g TradeGormRepo) GetTradesByUserIdAndExternalId(fields map[string]int) ([]domain.Trade, error) {
	var trades []domain.Trade
	tx := g.DB.Where(fields).Find(&trades)
	if tx.Error != nil {
		return []domain.Trade{}, tx.Error
	}
	return trades, nil
}
