package usecases

import (
	"cwapi/internal/core/domain"
	"cwapi/internal/core/ports"
	"cwapi/pkg/providers"
	"go.uber.org/zap"
)

type UserInteractor struct {
	userRepo        ports.UserRepository
	cacheRepo       ports.CacheRepository
	providerManager providers.ProviderManager
	logger          *zap.Logger
}

func NewUserInteractor(userRepo ports.UserRepository, cacheRepo ports.CacheRepository, providerManager *providers.ProviderManager) *UserInteractor {
	return &UserInteractor{userRepo: userRepo, cacheRepo: cacheRepo, providerManager: *providerManager}
}

func (interactor UserInteractor) Get(userID int) (domain.User, error) {
	return domain.NewUser(), nil
}

func (interactor UserInteractor) GetUserTrades(userID int) ([]domain.Trade, error) {
	//TODO implement me
	panic("implement me")
}
