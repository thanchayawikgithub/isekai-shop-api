package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/config"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases"
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

func (r *Router) RegisterRoutes() {
	r.registerItemShopRoutes()
	r.registerItemManagingRoutes()
	r.registerOAuth2Routes()
}
