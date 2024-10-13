package controllers

import "github.com/labstack/echo/v4"

type Oauth2Controller interface {
	PlayerLogin(ctx echo.Context) error
	AdminLogin(ctx echo.Context) error
	PlayerLoginCallback(ctx echo.Context) error
	AdminLoginCallback(ctx echo.Context) error
	Logout(ctx echo.Context) error
}
