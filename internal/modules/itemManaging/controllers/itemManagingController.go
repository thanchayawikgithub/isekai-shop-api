package controllers

import "github.com/labstack/echo/v4"

type ItemManagingController interface {
	Creating(ctx echo.Context) error
}
