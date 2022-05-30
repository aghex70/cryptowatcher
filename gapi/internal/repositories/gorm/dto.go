package gorm

import "gapi-agp/internal/core/domain"

func (u User) toDto() domain.User {
	return domain.User{
		ID:         u.ID,
		ExternalID: u.ExternalID,
		Source:     u.Source,
	}
}

func fromUserDto(u domain.User) User {
	return User{
		ID:         u.ID,
		ExternalID: u.ExternalID,
		Source:     u.Source,
	}
}

func (t Trade) toDto() domain.Trade {
	return domain.Trade{
		EventType:  t.EventType,
		ExternalID: int64(t.ExternalID),
		Symbol:     t.Symbol,
		Price:      t.Price,
		Currency:   t.Currency,
		Quantity:   t.Quantity,
		EventTime:  t.EventTime,
		TradeTime:  t.TradeTime,
		BuyerID:    int64(t.BuyerID),
		SellerID:   int64(t.SellerID),
		Source:     t.Source,
	}
}

func fromTradeDto(t domain.Trade) Trade {
	return Trade{
		EventType:  t.EventType,
		ExternalID: int(t.ExternalID),
		Symbol:     t.Symbol,
		Price:      t.Price,
		Currency:   t.Currency,
		Quantity:   t.Quantity,
		EventTime:  t.EventTime,
		TradeTime:  t.TradeTime,
		BuyerID:    int(t.BuyerID),
		SellerID:   int(t.SellerID),
		Source:     t.Source,
	}
}
