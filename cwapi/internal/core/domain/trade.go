package domain

import (
	"fmt"
	"time"

	"cwapi/internal/errors"
)

type TradeType string

const (
	Buy  TradeType = "BUY"
	Sell TradeType = "SELL"
)

type Trade struct {
	ProviderID  int64
	TradeID     int64
	SymbolsFrom string
	SymbolsTo   string
	Currency    string
	Price       float32
	Quantity    float32
	TradeType   TradeType
	EventTime   time.Time
	TradeTime   time.Time
	BuyerID     int64
	SellerID    int64
}

func NewTrade(maker bool, providerID int64, tradeID int64, symbolFrom string, symbolTo string, currency string, price float32, quantity float32, eventTime time.Time, tradeTime time.Time, buyerID int64, sellerID int64) Trade {
	return Trade{
		ProviderID:  providerID,
		TradeID:     tradeID,
		SymbolsFrom: symbolFrom,
		SymbolsTo:   symbolTo,
		Currency:    currency,
		Price:       price,
		Quantity:    quantity,
		TradeType:   MakerToTradeType(maker),
		EventTime:   eventTime,
		TradeTime:   tradeTime,
		BuyerID:     buyerID,
		SellerID:    sellerID,
	}
}

func MakerToTradeType(maker bool) TradeType {
	if maker == true {
		return Sell
	}
	return Buy
}

func (t Trade) IsValid() error {
	if t.ValidateUsers() == false {
		fmt.Println("Invalid trade: Buyer and Seller are the same")
		return errors.ErrInvalidUsers
	}
	return nil
}

func (t Trade) ValidateUsers() bool {
	if t.BuyerID == t.SellerID {
		return false
	}
	return true
}
