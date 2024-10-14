package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/config"
	oauth2Controllers "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/oauth2/controllers"
)

type authMiddleWare struct {
	oauth2Controller oauth2Controllers.Oauth2Controller
	oauth2Config     *config.OAuth2
	logger           echo.Logger
}

func NewAuthMiddleware(oauth2Controller oauth2Controllers.Oauth2Controller,
	oauth2Config *config.OAuth2,
	logger echo.Logger) *authMiddleWare {
	return &authMiddleWare{oauth2Controller, oauth2Config, logger}
}

func (mw *authMiddleWare) PlayerAuthorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return mw.oauth2Controller.PlayerAuthorize(ctx, next)
	}
}

func (mw *authMiddleWare) AdminAuthorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return mw.oauth2Controller.AdminAuthorize(ctx, next)
	}
}
