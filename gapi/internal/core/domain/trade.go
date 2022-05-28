package domain

import (
	"log"
	"time"

	"gapi-agp/internal/errors"
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
		log.Logger.Println("Invalid trade: Buyer and Seller are the same")
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

//{"e":"trade","E":1653432221597,"s":"BTCUSDT","t":1380155726,"p":"29561.44000000","q":"0.01691000","b":10714035774,"a":10714036214,"T":1653432221597,"m":true,"M":true}

//{
//"e": "trade",     // Event type
//"E": 123456789,   // Event time
//"s": "BNBBTC",    // Symbol
//"t": 12345,       // Trade ID
//"p": "0.001",     // Price
//"q": "100",       // Quantity
//"b": 88,          // Buyer order ID
//"a": 50,          // Seller order ID
//"T": 123456785,   // Trade time
//"m": true,        // Is the buyer the market maker?  FALSE -> BUY, TRUE -> SELL
//"M": true         // Ignore
//}
