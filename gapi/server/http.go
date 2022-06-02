package server

import (
	"fmt"
	"gapi-agp/config"
	"gapi-agp/internal/core/handlers"
	"gapi-agp/internal/core/ports"
	"net/http"
)

type Server struct {
	handler      handlers.RestHandler
	tradeUseCase ports.TradeUseCase
	userUseCase  ports.UserUseCase
}

func (s *Server) StartServer() error {

	handler := handlers.NewRestHandler(s.tradeUseCase, s.userUseCase)

	http.HandleFunc("/fetch", handler.FetchOrders)
	http.HandleFunc("/stop", handler.StopFetchOrders)
	http.HandleFunc("/trades", handler.GetTrades)
	http.HandleFunc("/sales", handler.GetSales)
	http.HandleFunc("/purchases", handler.GetPurchases)
	http.HandleFunc("/user", handler.GetUser)
	http.HandleFunc("/user/trades", handler.GetUserTrades)

	address := fmt.Sprintf("%s:%d", config.C.Server.Host, config.C.Server.Port)
	fmt.Println("Server started at", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}
	return nil
}

func NewServer(tradeUseCase ports.TradeUseCase, userUseCase ports.UserUseCase) *Server {
	return &Server{
		tradeUseCase: tradeUseCase,
		userUseCase:  userUseCase,
	}
}
