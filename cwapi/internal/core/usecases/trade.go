package usecases

import (
	"context"
	"cwapi/config"
	"cwapi/internal/core/domain"
	"cwapi/internal/core/ports"
	"cwapi/pkg/providers"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type TradeInteractor struct {
	tradeRepo       ports.TradeRepository
	cacheRepo       ports.CacheRepository
	providerManager providers.ProviderManager
	logger          *zap.Logger
}

func NewTradeInteractor(tradeRepo ports.TradeRepository, cacheRepo ports.CacheRepository, providerManager *providers.ProviderManager, logger *zap.Logger) *TradeInteractor {
	return &TradeInteractor{tradeRepo: tradeRepo, cacheRepo: cacheRepo, providerManager: *providerManager, logger: logger}
}

func (interactor TradeInteractor) Fetch(r ports.FetchRequest) (*providers.FetchTradesResponse, error) {
	interactor.logger.Info("Starting fetching process in provider")
	address := fmt.Sprintf("%s:%d", config.C.Providers.Fetcher.Host, config.C.Providers.Fetcher.Port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		interactor.logger.Error("failed to create gRPC channel", zap.Error(err))
		return &providers.FetchTradesResponse{}, err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			interactor.logger.Fatal("error closing gRPC channel", zap.Error(err))
		}
	}(conn)

	client := providers.NewFetcherServiceClient(conn)
	res, err := client.FetchTrades(context.Background(), &providers.FetchTradesRequest{Source: r.Source})
	if err != nil {
		interactor.logger.Error("error starting fetching process in provider", zap.Error(err))
		return &providers.FetchTradesResponse{}, err
	}
	interactor.logger.Info("Started fetching process in provider")
	return res, nil
}

func (interactor TradeInteractor) StopFetch(r ports.StopFetchRequest) error {
	interactor.logger.Info("Ending fetching process in provider")
	address := fmt.Sprintf("%s:%d", config.C.Providers.Fetcher.Host, config.C.Providers.Fetcher.Port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		interactor.logger.Error("failed to create gRPC channel", zap.Error(err))
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			interactor.logger.Fatal("error closing gRPC channel", zap.Error(err))
		}
	}(conn)

	client := providers.NewFetcherServiceClient(conn)
	_, err = client.StopFetchTrades(context.Background(), &providers.Empty{})
	if err != nil {
		interactor.logger.Error("error ending fetching process in provider", zap.Error(err))
		return err
	}
	interactor.logger.Info("Ended fetching process in provider")
	return nil
}

func (interactor TradeInteractor) Get() ([]domain.Trade, error) {
	interactor.logger.Info("Retrieving trades")
	trades, err := interactor.tradeRepo.GetTrades()
	if err != nil {
		return []domain.Trade{}, err
	}
	return trades, nil
}

func (interactor TradeInteractor) GetSales() ([]domain.Trade, error) {
	interactor.logger.Info("Retrieving sales")
	trades, err := interactor.tradeRepo.GetTradesByEventType()
	if err != nil {
		return []domain.Trade{}, err
	}
	return trades, nil
}

func (interactor TradeInteractor) GetPurchases() ([]domain.Trade, error) {
	interactor.logger.Info("Retrieving purchases")
	trades, err := interactor.tradeRepo.GetTrades()
	if err != nil {
		return []domain.Trade{}, err
	}
	return trades, nil
}
