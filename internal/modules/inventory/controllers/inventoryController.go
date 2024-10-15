package controllers

import "github.com/labstack/echo/v4"

type InventoryController interface {
	Listing(ctx echo.Context) error
}
