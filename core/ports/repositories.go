package ports

import (
	"context"
	"crud/core/domain"
)

type Persistence interface {
	Create(ctx context.Context, input *domain.CreateUserInput) (*domain.CreateUserOutput, error)
	Read(ctx context.Context, input *domain.ReadUserInput) (*domain.ReadUserOutput, error)
	Update(ctx context.Context, input *domain.UpdateUserInput) (*domain.UpdateUserOutput, error)
	Delete(ctx context.Context, input *domain.DeleteUserInput) (*domain.DeleteUserOutput, error)
}
