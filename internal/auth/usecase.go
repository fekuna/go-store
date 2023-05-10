package auth

import (
	"context"

	"github.com/fekuna/go-store/internal/models"
)

type UseCase interface {
	Register(ctx context.Context, user *models.User) (*models.UserWithToken, error)
}
