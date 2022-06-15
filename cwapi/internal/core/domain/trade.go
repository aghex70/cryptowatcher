package domain

import (
	"fmt"
	"time"

	"cwapi/internal/errors"
)

type EventType string

const (
	Sell EventType = "sell"
	Buy  EventType = "buy"
)

type Trade struct {
	EventType  EventType
	ExternalID int64
	Symbol     string
	Price      float32
	Currency   string
	Quantity   float32
	EventTime  time.Time
	TradeTime  time.Time
	BuyerID    int64
	SellerID   int64
	Source     string
}

func NewTrade(maker bool, externalID int64, symbol string, price float32, quantity float32, eventTime time.Time, tradeTime time.Time, buyerID int64, sellerID int64, source string) Trade {
	return Trade{
		EventType:  MakerToEventType(maker),
		ExternalID: externalID,
		Symbol:     symbol,
		Price:      price,
		Quantity:   quantity,
		EventTime:  eventTime,
		TradeTime:  tradeTime,
		BuyerID:    buyerID,
		SellerID:   sellerID,
		Source:     source,
	}
}

func MakerToEventType(maker bool) EventType {
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
