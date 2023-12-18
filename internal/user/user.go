package user

import (
	"context"

	"github.com/squacky/openchat/internal/user/domain"
)

type service struct {
	repository Repository
}

func (s *service) CreateUser(ctx context.Context, user *domain.User) error {
	return s.repository.CreateUser(ctx, user)
}

func (s *service) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	return s.repository.GetUserByID(ctx, id)
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
