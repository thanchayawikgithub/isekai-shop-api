package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/config"
)

type middleware struct {
	app  *echo.Echo
	conf *config.Server
}

// NewMiddleware creates a new middleware instance.
func NewMiddleware(app *echo.Echo, conf *config.Server) *middleware {
	return &middleware{app, conf}
}

// RegisterMiddleWares registers the application middlewares.
func (m *middleware) RegisterMiddleWares() {
	m.app.Use(echoMiddleware.Recover())
	m.app.Use(echoMiddleware.Logger())
	m.app.Use(corsMiddleware(m.conf.AllowOrigins))
	m.app.Use(bodyLimitMiddleware(m.conf.BodyLimit))
	m.app.Use(timeOutMiddleware(m.conf.Timeout))
}

// TimeOutMiddleware returns a middleware that times out requests.
func timeOutMiddleware(timeoutDuration time.Duration) echo.MiddlewareFunc {
	print(timeoutDuration)
	return echoMiddleware.TimeoutWithConfig(echoMiddleware.TimeoutConfig{
		Skipper:      echoMiddleware.DefaultSkipper,
		ErrorMessage: "Request Timeout",
		Timeout:      timeoutDuration * time.Second,
	})
}

// CORS Middleware to handle Cross-Origin Resource Sharing
func corsMiddleware(allowOrigins []string) echo.MiddlewareFunc {
	return echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		Skipper:      echoMiddleware.DefaultSkipper,
		AllowOrigins: allowOrigins,
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})
}

// BodyLimitMiddleware restricts the size of request bodies.
func bodyLimitMiddleware(bodyLimit string) echo.MiddlewareFunc {
	return echoMiddleware.BodyLimit(bodyLimit)
}
