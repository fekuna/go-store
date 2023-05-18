package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/fekuna/go-store/config"
	"github.com/fekuna/go-store/internal/auth"
	"github.com/fekuna/go-store/internal/models"
	"github.com/fekuna/go-store/internal/session"
	"github.com/fekuna/go-store/pkg/httpErrors"
	"github.com/fekuna/go-store/pkg/logger"
	"github.com/fekuna/go-store/pkg/utils"
	"github.com/labstack/echo/v4"
)

// Auth handlers
type authHandlers struct {
	cfg    *config.Config
	logger logger.Logger
	authUC auth.UseCase
	sessUC session.UseCase
}

func NewAuthHandlers(cfg *config.Config, logger logger.Logger, authUC auth.UseCase, sessUC session.UseCase) auth.Handlers {
	return &authHandlers{
		cfg:    cfg,
		logger: logger,
		authUC: authUC,
		sessUC: sessUC,
	}
}

// Register godoc
// @Summary Register new user
// @Description register new user, returns user and token
// @Tags Auth
// @Accept json
// @Produce json
// @Success 201 {object} models.User
// @Router /auth/register [post]
func (h *authHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: tracing
		ctx := context.Background()

		user := &models.User{}
		if err := utils.ReadRequest(c, user); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		createdUser, err := h.authUC.Register(ctx, user)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		session := &models.Session{
			RefreshToken: createdUser.Token.RefreshToken,
			ExpiresAt:    time.Now().Add(time.Hour * 24 * 30),
			UserID:       createdUser.User.UserID,
		}

		_, err = h.sessUC.CreateSession(ctx, session)
		if err != nil {
			fmt.Println("owkowkowk", err.Error())
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, createdUser)
	}
}

// Login godoc
// @Summary Login user
// @Description login user, returns tokens and set session in DB
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /auth/login [post]
func (h *authHandlers) Login() echo.HandlerFunc {
	type Login struct {
		Email    string `json:"email" db:"email" validate:"omitempty,lte=60"`
		Password string `json:"password,omitempty" db:"password" validate:"required,gte=6"`
	}
	return func(c echo.Context) error {
		// TODO: tracing
		ctx := context.Background()

		login := &Login{}
		if err := utils.ReadRequest(c, login); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		userWithToken, err := h.authUC.Login(ctx, &models.User{
			Email:    login.Email,
			Password: login.Password,
		})
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		sess := &models.Session{
			RefreshToken: userWithToken.Token.RefreshToken,
			ExpiresAt:    time.Now().Add(time.Hour * 24 * 30),
			UserID:       userWithToken.User.UserID,
		}

		// TODO: can generate multiple session for multiple devices. for the future feature.
		_, err = h.sessUC.UpsertSession(ctx, sess)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		// fmt.Println(upserted)

		return c.JSON(http.StatusOK, userWithToken)

		// sess, err := h.sessUC.
	}
}
