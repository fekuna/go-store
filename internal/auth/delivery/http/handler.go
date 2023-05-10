package http

import (
	"context"
	"net/http"

	"github.com/fekuna/go-store/config"
	"github.com/fekuna/go-store/internal/auth"
	"github.com/fekuna/go-store/internal/models"
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
}

func NewAuthHandlers(cfg *config.Config, logger logger.Logger, authUC auth.UseCase) auth.Handlers {
	return &authHandlers{
		cfg:    cfg,
		logger: logger,
		authUC: authUC,
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

		return c.JSON(http.StatusCreated, createdUser)
	}
}
