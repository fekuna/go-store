package http

import (
	"github.com/fekuna/go-store/internal/auth"
	"github.com/fekuna/go-store/pkg/middleware"
	"github.com/labstack/echo/v4"
)

func MapAuthRoutes(authGroup *echo.Group, h auth.Handlers, mw *middleware.MiddlewareManager) {
	authGroup.POST("/register", h.Register())
	authGroup.POST("/login", h.Login())
}
