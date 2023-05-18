package usecase

import (
	"context"
	"net/http"
	"time"

	"github.com/fekuna/go-store/config"
	"github.com/fekuna/go-store/internal/auth"
	"github.com/fekuna/go-store/internal/models"
	"github.com/fekuna/go-store/pkg/httpErrors"
	"github.com/fekuna/go-store/pkg/logger"
	"github.com/fekuna/go-store/pkg/utils"
	"github.com/pkg/errors"
)

// Auth Usecase
type authUC struct {
	cfg      *config.Config
	logger   logger.Logger
	authRepo auth.Repository
}

// Auth usecase constructor
func NewAuthUseCase(cfg *config.Config, logger logger.Logger, authRepo auth.Repository) *authUC {
	return &authUC{
		cfg:      cfg,
		logger:   logger,
		authRepo: authRepo,
	}
}

func (u *authUC) Register(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	// TODO: Tracing

	existsUser, err := u.authRepo.FindByEmail(ctx, user)
	if existsUser != nil || err == nil {
		return nil, httpErrors.NewRestErrorWithMessage(http.StatusBadRequest, httpErrors.ErrEmailAlreadyExists, err)
	}

	if err = user.PrepareCreate(); err != nil {
		return nil, httpErrors.NewBadRequestError(errors.Wrap(err, "authUC.Register.PrepareCreate"))
	}

	createdUser, err := u.authRepo.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	createdUser.SanitizePassword()

	accessToken, err := utils.GenerateJWTToken(createdUser, u.cfg, time.Minute*30)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "authUC.Register.AccessToken.GenerateJWTTOken"))
	}

	refreshToken, err := utils.GenerateJWTToken(createdUser, u.cfg, (time.Hour*24)*30)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "authUC.Register.AccessToken.GenerateJWTTOken"))
	}

	authToken := models.AuthToken{
		AccesToken:   accessToken,
		RefreshToken: refreshToken,
	}

	return &models.UserWithToken{
		User:  createdUser,
		Token: authToken,
	}, nil
}
