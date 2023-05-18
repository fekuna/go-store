package repository

import (
	"context"

	"github.com/fekuna/go-store/internal/models"
	"github.com/fekuna/go-store/internal/session"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Session repository
type sessionRepo struct {
	db *sqlx.DB
}

// Session New constructor
func NewSessionRepository(db *sqlx.DB) session.Repository {
	return &sessionRepo{
		db: db,
	}
}

func (r *sessionRepo) CreateSession(ctx context.Context, sess *models.Session) (*models.Session, error) {
	// TODO: tracing

	s := &models.Session{}
	if err := r.db.QueryRowxContext(
		ctx, createSession, &sess.RefreshToken, &sess.ExpiresAt, &sess.UserID,
	).StructScan(s); err != nil {
		return nil, errors.Wrap(err, "sessionRepo.CreateSession.StructScan")
	}

	return s, nil
}
