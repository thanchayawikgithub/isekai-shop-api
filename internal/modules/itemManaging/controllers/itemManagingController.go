package controllers

import "github.com/labstack/echo/v4"

type ItemManagingController interface {
	Creating(ctx echo.Context) error
	Editing(ctx echo.Context) error
	Archiving(ctx echo.Context) error
}
