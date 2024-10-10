package controllers

import "github.com/labstack/echo/v4"

type ItemShopController interface {
	Listing(ctx echo.Context) error
}
