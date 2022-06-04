package usecases

import (
	"context"
	"fmt"
	"gapi-agp/config"
	"gapi-agp/internal/core/domain"
	"gapi-agp/internal/core/ports"
	"gapi-agp/internal/core/usecases/pb"
	"gapi-agp/pkg/providers"
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

func (interactor TradeInteractor) Fetch(r ports.FetchRequest) error {
	interactor.logger.Info("Starting fetching process in provider")
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

	client := pb.NewFetcherServiceClient(conn)
	interactor.logger.Info("Creating fetcher gRPC client")
	source := pb.DataSource{Source: r.Source}
	_, err = client.FetchTrades(context.Background(), &pb.FetchTradesRequest{Sources: []*pb.DataSource{&source}})
	if err != nil {
		interactor.logger.Error("error starting fetching process in provider", zap.Error(err))
		return err
	}
	return nil
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

	client := pb.NewFetcherServiceClient(conn)
	interactor.logger.Info("Creating fetcher gRPC client")
	_, err = client.StopFetchTrades(context.Background(), &pb.Empty{})
	if err != nil {
		interactor.logger.Error("error ending fetching process in provider", zap.Error(err))
		return err
	}
	return nil
}

func (interactor TradeInteractor) Get() ([]domain.Trade, error) {
	interactor.logger.Info("Retrieving trades")
	return []domain.Trade{}, nil
}

func (interactor TradeInteractor) GetSales() ([]domain.Trade, error) {
	interactor.logger.Info("Retrieving sales")
	return []domain.Trade{}, nil
}

func (interactor TradeInteractor) GetPurchases() ([]domain.Trade, error) {
	interactor.logger.Info("Retrieving purchases")
	return []domain.Trade{}, nil
}

//type FetchOrdersRequest struct {
//	Source string
//}
//
//type StopFetchOrdersRequest struct {
//	Source string
//}
