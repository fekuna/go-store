package repository

import (
	"context"

	"github.com/fekuna/go-store/internal/auth"
	"github.com/fekuna/go-store/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) auth.Repository {
	return &authRepo{
		db: db,
	}
}

func (r *authRepo) Register(ctx context.Context, user *models.User) (*models.User, error) {
	// TODO: Tracing
	u := &models.User{}
	if err := r.db.QueryRowxContext(
		ctx, createUserQuery, &user.FirstName, &user.LastName, &user.Email,
		&user.Password, &user.Role, &user.About, &user.Avatar, &user.PhoneNumber, &user.Address, &user.City,
		&user.Gender, &user.Postcode, &user.Birthday,
	).StructScan(u); err != nil {
		return nil, errors.Wrap(err, "authRepo.Register.StructScan")
	}

	return u, nil
}

func (r *authRepo) FindByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	// TODO: findByEmail

	foundUser := &models.User{}
	if err := r.db.QueryRowxContext(ctx, findUserByEmail, user.Email).StructScan(foundUser); err != nil {
		return nil, errors.Wrap(err, "authRepo.FindByEmail.QueryRowxContext")
	}

	return foundUser, nil
}
