package usecase

import (
	"context"

	"github.com/fekuna/go-store/config"
	"github.com/fekuna/go-store/internal/models"
	"github.com/fekuna/go-store/internal/session"
	"github.com/fekuna/go-store/pkg/logger"
)

type SessionUC struct {
	cfg         *config.Config
	logger      logger.Logger
	sessionRepo session.Repository
}

func NewSessionUseCase(cfg *config.Config, logger logger.Logger, sessionRepo session.Repository) session.UseCase {
	return &SessionUC{
		cfg:         cfg,
		logger:      logger,
		sessionRepo: sessionRepo,
	}
}

func (s *SessionUC) CreateSession(ctx context.Context, session *models.Session) (*models.Session, error) {
	// TODO: Tracing

	return s.sessionRepo.CreateSession(ctx, session)
}
