package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (s *SessionUC) UpsertSession(ctx context.Context, session *models.Session) (*models.Session, error) {
	// TODO: tracing

	_, err := s.sessionRepo.FindSessionByUserId(ctx, session)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	// If No Data found. we insert it
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Create session")
		return s.sessionRepo.CreateSession(ctx, session)
	} else {
		fmt.Println("Update session")
		// update session if user has session
		return s.sessionRepo.UpdateSessionByUserId(ctx, session)
	}
}
