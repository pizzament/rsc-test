package service

import (
	"context"
	"fmt"
	"time"

	"github.com/pizzament/rsc-test/internal/model"
)

type repository interface {
	AddCount(ctx context.Context, bannerID model.BannerID, timestamp time.Time) error
	ReceiveStats(ctx context.Context, bannerID model.BannerID, from time.Time, to time.Time) ([]model.Stat, error)
}

type Service struct {
	repository repository
}

// NewService инициализируем сервис.
func NewService(repository repository) *Service {
	return &Service{repository: repository}
}

// AddCount метод сервиса для увеличения счётчика для bannerID.
func (s *Service) AddCount(ctx context.Context, bannerID model.BannerID) error {
	now := time.Now()
	truncatedNow := now.Truncate(time.Minute)

	err := s.repository.AddCount(ctx, bannerID, truncatedNow)
	if err != nil {
		return fmt.Errorf("Service.AddCount error: %w", err)
	}

	return nil
}

// ReceiveStats метод сервиса для получения данных по bannerID.
func (s *Service) ReceiveStats(ctx context.Context, bannerID model.BannerID, from time.Time, to time.Time) ([]model.Stat, error) {
	stats, err := s.repository.ReceiveStats(ctx, bannerID, from, to)
	if err != nil {
		return nil, fmt.Errorf("Service.ReceiveStats error: %w", err)
	}

	return stats, nil
}
