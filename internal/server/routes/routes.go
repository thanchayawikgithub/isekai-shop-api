package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Router struct {
	app    *echo.Echo
	db     *gorm.DB
	logger echo.Logger
}

func NewRouter(app *echo.Echo, db *gorm.DB, logger echo.Logger) *Router {
	return &Router{app, db, logger}
}

func (r *Router) RegisterRoutes() {
	r.registerItemShopRoutes()
}
