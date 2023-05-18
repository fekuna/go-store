package session

import (
	"context"

	"github.com/fekuna/go-store/internal/models"
)

type UseCase interface {
	CreateSession(ctx context.Context, session *models.Session) (*models.Session, error)
}
