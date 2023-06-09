package auth

import (
	"context"
	"net/url"

	"github.com/fekuna/go-store/internal/models"
	"github.com/google/uuid"
)

type UseCase interface {
	Register(ctx context.Context, user *models.User) (*models.UserWithToken, error)
	Login(ctx context.Context, user *models.User) (*models.UserWithToken, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	UploadAvatar(ctx context.Context, userID uuid.UUID, file models.UploadInput) (*models.User, error)
	GetAvatar(ctx context.Context) (*url.URL, error)
}
