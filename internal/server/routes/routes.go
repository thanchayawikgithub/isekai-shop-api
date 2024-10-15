package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/config"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/server/middlewares"
)

type Router struct {
	app    *echo.Echo
	db     databases.Database
	logger echo.Logger
	config *config.Config
}

func NewRouter(app *echo.Echo, db databases.Database, logger echo.Logger, config *config.Config) *Router {
	return &Router{app, db, logger, config}
}

func (r *Router) RegisterRoutes(authMiddleWare *middlewares.AuthMiddleWare) {
	r.registerItemShopRoutes()
	r.registerItemManagingRoutes(authMiddleWare)
	r.registerOAuth2Routes()
	r.registerPlayerCoinRoutes(authMiddleWare)
}
