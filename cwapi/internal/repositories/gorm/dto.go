package gorm

import "cwapi/internal/core/domain"

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
		ProviderID:  int64(t.ProviderID),
		TradeID:     int64(t.TradeID),
		SymbolsFrom: t.SymbolsFrom,
		SymbolsTo:   t.SymbolsTo,
		Currency:    t.Currency,
		Price:       t.Price,
		Quantity:    t.Quantity,
		TradeType:   t.TradeType,
		EventTime:   t.EventTime,
		TradeTime:   t.TradeTime,
		BuyerID:     int64(t.BuyerID),
		SellerID:    int64(t.SellerID),
	}
}

func fromTradeDto(t domain.Trade) Trade {
	return Trade{
		ProviderID:  int(t.ProviderID),
		TradeID:     int(t.TradeID),
		SymbolsFrom: t.SymbolsFrom,
		SymbolsTo:   t.SymbolsTo,
		Currency:    t.Currency,
		Price:       t.Price,
		Quantity:    t.Quantity,
		TradeType:   t.TradeType,
		EventTime:   t.EventTime,
		TradeTime:   t.TradeTime,
		BuyerID:     int(t.BuyerID),
		SellerID:    int(t.SellerID),
	}
}
