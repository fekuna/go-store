package http

import (
	"github.com/fekuna/go-store/internal/auth"
	"github.com/fekuna/go-store/internal/middleware"
	"github.com/labstack/echo/v4"
)

func MapAuthRoutes(authGroup *echo.Group, h auth.Handlers, mw *middleware.MiddlewareManager) {
	authGroup.POST("/register", h.Register())
	authGroup.POST("/login", h.Login())
	authGroup.Use(mw.AuthJWTMiddleware)
	authGroup.GET("/me", h.GetMe())
	authGroup.POST("/:user_id/avatar", h.UploadAvatar())
	authGroup.GET("/avatar", h.GetAvatar())
}
