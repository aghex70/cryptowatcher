package handlers

import (
	"encoding/json"
	"gapi-agp/internal/core/ports"
	"net/http"
)

type RestHandler struct {
	tradeUseCase ports.TradeUseCase
	userUseCase  ports.UserUseCase
}

func NewRestHandler(tradeUseCase ports.TradeUseCase, userUseCase ports.UserUseCase) RestHandler {
	return RestHandler{
		tradeUseCase: tradeUseCase,
		userUseCase:  userUseCase,
	}
}

func (h RestHandler) FetchOrders(w http.ResponseWriter, r *http.Request) {
	// Get source param from request body
	source := "binance"
	err := h.tradeUseCase.Fetch(source)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h RestHandler) StopFetchOrders(w http.ResponseWriter, r *http.Request) {
	source := "binance"
	err := h.tradeUseCase.StopFetch(source)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h RestHandler) GetTrades(w http.ResponseWriter, r *http.Request) {
	trades, err := h.tradeUseCase.Get()
	b, err := json.Marshal(trades)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h RestHandler) GetSales(w http.ResponseWriter, r *http.Request) {
	sales, err := h.tradeUseCase.GetSales()
	b, err := json.Marshal(sales)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h RestHandler) GetPurchases(w http.ResponseWriter, r *http.Request) {
	purchases, err := h.tradeUseCase.GetPurchases()
	b, err := json.Marshal(purchases)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h RestHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := 1
	user, err := h.userUseCase.Get(userID)
	b, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h RestHandler) GetUserTrades(w http.ResponseWriter, r *http.Request) {
	userID := 1
	trades, err := h.userUseCase.GetUserTrades(userID)
	b, err := json.Marshal(trades)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
