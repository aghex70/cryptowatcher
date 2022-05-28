package usecases

import (
	"gapi-agp/internal/core/domain"
	"gapi-agp/internal/core/ports"
)

type UserInteractor struct {
	userRepo ports.UserRepository
}

func (interactor UserInteractor) Get(userID int) (domain.User, error) {
	return domain.NewUser(), nil
}

func (interactor UserInteractor) GetTrades(userID int) ([]domain.Trade, error) {
	return []domain.Trade, nil
}