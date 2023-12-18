package user

import (
	"context"

	"github.com/squacky/openchat/internal/user/domain"
)

type Service interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
}
