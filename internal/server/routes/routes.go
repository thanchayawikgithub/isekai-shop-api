package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases"
)

type Router struct {
	app    *echo.Echo
	db     databases.Database
	logger echo.Logger
}

func NewRouter(app *echo.Echo, db databases.Database, logger echo.Logger) *Router {
	return &Router{app, db, logger}
}

func (r *Router) RegisterRoutes() {
	r.registerItemShopRoutes()
	r.registerItemManagingRoutes()
}
