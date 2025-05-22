package service

import (
	"context"

	"github.com/pizzament/rsc-test/internal/model"
)

type Repository interface {
}

type Service struct {
	Repository Repository
}

// NewService инициализируем сервис
func NewService(Repository Repository) *Service {
	return &Service{Repository: Repository}
}

// AddCount метод увеличения счётчика для bannerID
func (s *Service) AddCount(ctx context.Context, bannerID model.BannerID) error {
	return nil
}

// ReceiveStats метод для получения данных по bannerID
func (s *Service) ReceiveStats(ctx context.Context, bannerID model.BannerID) ([]model.Stat, error) {
	return nil, nil
}
