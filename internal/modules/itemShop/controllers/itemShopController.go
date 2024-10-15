package controllers

import "github.com/labstack/echo/v4"

type ItemShopController interface {
	Listing(ctx echo.Context) error
	Buying(ctx echo.Context) error
	Selling(ctx echo.Context) error
}
