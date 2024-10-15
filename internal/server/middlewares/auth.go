package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/config"
	oauth2Controllers "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/oauth2/controllers"
)

type AuthMiddleWare struct {
	oauth2Controller oauth2Controllers.Oauth2Controller
	oauth2Config     *config.OAuth2
	logger           echo.Logger
}

func NewAuthMiddleware(oauth2Controller oauth2Controllers.Oauth2Controller,
	oauth2Config *config.OAuth2,
	logger echo.Logger) *AuthMiddleWare {
	return &AuthMiddleWare{oauth2Controller, oauth2Config, logger}
}

func (mw *AuthMiddleWare) PlayerAuthorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return mw.oauth2Controller.PlayerAuthorize(ctx, next)
	}
}

func (mw *AuthMiddleWare) AdminAuthorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return mw.oauth2Controller.AdminAuthorize(ctx, next)
	}
}
