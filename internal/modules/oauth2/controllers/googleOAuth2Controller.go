package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/config"
	oauth2Service "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/oauth2/services"
)

type googleOAuth2Controller struct {
	oauth2Service oauth2Service.Oauth2Service
	oauth2Conf    *config.OAuth2
	logger        echo.Logger
}

func NewGoogleOAuth2Service(oauth2Service oauth2Service.Oauth2Service,
	oauth2Conf *config.OAuth2,
	logger echo.Logger) oauth2Service.Oauth2Service {
	return &googleOAuth2Controller{oauth2Service, oauth2Conf, logger}
}
