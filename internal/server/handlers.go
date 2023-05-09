package server

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	apiMiddlewares "github.com/fekuna/go-store/pkg/middleware"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init Repository

	// Init useCase

	// Init handlers

	mw := apiMiddlewares.NewMiddlewareManager(s.cfg, s.logger)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))
	if s.cfg.Server.Debug {
		e.Use(mw.DebugMiddleware)
	}

	return nil
}
