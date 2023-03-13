package ports

import (
	"context"
	"crud/core/domain"
)

type CreateQrCode interface {
	CreateQrInfo(ctx context.Context, input *domain.UserQrCodeInput) (*domain.UserQrCodeOutput, error)
	Approve(ctx context.Context, input *domain.ApproveInput) (*domain.ApproveOutput, error)
}
