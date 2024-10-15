package controllers

import "github.com/labstack/echo/v4"

type PlayerCoinController interface {
	CoinAdding(ctx echo.Context) error
	CoinShowing(ctx echo.Context) error
}
