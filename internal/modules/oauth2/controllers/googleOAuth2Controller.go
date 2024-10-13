package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/avast/retry-go"
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/config"
	adminModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/admin/models"
	oauth2Exceptions "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/oauth2/exceptions"
	oauth2Models "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/oauth2/models"
	oauth2Services "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/oauth2/services"
	playerModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/player/models"
	"github.com/thanchayawikgithub/isekai-shop-api/pkg/custom"
	"golang.org/x/exp/rand"
	"golang.org/x/oauth2"
)

type googleOAuth2Controller struct {
	oauth2Service oauth2Services.Oauth2Service
	oauth2Conf    *config.OAuth2
	logger        echo.Logger
}

var (
	playerGoogleOAuth2 *oauth2.Config
	adminGoogleOAuth2  *oauth2.Config
	once               sync.Once

	accessTokenCookieName  = "act"
	refreshTokenCookieName = "rft"
	stateCookieName        = "state"

	letterBytes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func NewGoogleOAuth2Controller(oauth2Service oauth2Services.Oauth2Service,
	oauth2Conf *config.OAuth2,
	logger echo.Logger) Oauth2Controller {
	once.Do(func() {
		setGoogleOAuth2Config(oauth2Conf)
	})

	return &googleOAuth2Controller{oauth2Service, oauth2Conf, logger}
}

func setGoogleOAuth2Config(oauth2Conf *config.OAuth2) {
	playerGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientId,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.PlayerRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}

	adminGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientId,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.AdminRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}
}

func (c *googleOAuth2Controller) PlayerLogin(ctx echo.Context) error {
	state := c.randomState()

	c.setCookie(ctx, stateCookieName, state)

	return ctx.Redirect(http.StatusFound, playerGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) AdminLogin(ctx echo.Context) error {
	state := c.randomState()

	c.setCookie(ctx, stateCookieName, state)

	return ctx.Redirect(http.StatusFound, adminGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) PlayerLoginCallback(ctx echo.Context) error {
	context := context.Background()

	if err := retry.Do(func() error {
		return c.callbackValidate(ctx)
	}, retry.Attempts(3), retry.Delay(3*time.Second)); err != nil {
		c.logger.Errorf("Failed to validate callback: %s", err.Error())
		return custom.Error(ctx, http.StatusUnauthorized, err)
	}

	token, err := playerGoogleOAuth2.Exchange(context, ctx.QueryParam("code"))
	if err != nil {
		c.logger.Errorf("Failed to exchange token: %s", err.Error())
		return custom.Error(ctx, http.StatusUnauthorized, &oauth2Exceptions.Unauthorized{})
	}

	client := playerGoogleOAuth2.Client(context, token)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Failed to get user info: %s", err.Error())
		return custom.Error(ctx, http.StatusUnauthorized, &oauth2Exceptions.Unauthorized{})
	}

	playerCreatingReq := &playerModels.PlayerCreatingReq{
		ID:     userInfo.ID,
		Email:  userInfo.Email,
		Name:   userInfo.Name,
		Avatar: userInfo.Picture,
	}

	if err := c.oauth2Service.PlayerAccountCreating(playerCreatingReq); err != nil {
		c.logger.Errorf("Failed to create player account: %s", err.Error())
		return custom.Error(ctx, http.StatusInternalServerError, &oauth2Exceptions.OAuth2Processing{})
	}

	c.setSameSiteCookie(ctx, accessTokenCookieName, token.AccessToken)
	c.setSameSiteCookie(ctx, refreshTokenCookieName, token.RefreshToken)

	return ctx.JSON(http.StatusOK, &oauth2Models.LoginResponse{Message: "Login Success"})
}

func (c *googleOAuth2Controller) AdminLoginCallback(ctx echo.Context) error {
	context := context.Background()

	if err := retry.Do(func() error {
		return c.callbackValidate(ctx)
	}, retry.Attempts(3), retry.Delay(3*time.Second)); err != nil {
		c.logger.Errorf("Failed to validate callback: %s", err.Error())
		return custom.Error(ctx, http.StatusUnauthorized, err)
	}

	token, err := adminGoogleOAuth2.Exchange(context, ctx.QueryParam("code"))
	if err != nil {
		c.logger.Errorf("Failed to exchange token: %s", err.Error())
		return custom.Error(ctx, http.StatusUnauthorized, &oauth2Exceptions.Unauthorized{})
	}

	client := adminGoogleOAuth2.Client(context, token)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Failed to get user info: %s", err.Error())
		return custom.Error(ctx, http.StatusUnauthorized, &oauth2Exceptions.Unauthorized{})
	}

	adminCreatingReq := &adminModels.AdminCreatingReq{
		ID:     userInfo.ID,
		Email:  userInfo.Email,
		Name:   userInfo.Name,
		Avatar: userInfo.Picture,
	}

	if err := c.oauth2Service.AdminAccountCreating(adminCreatingReq); err != nil {
		c.logger.Errorf("Failed to create admin account: %s", err.Error())
		return custom.Error(ctx, http.StatusInternalServerError, &oauth2Exceptions.OAuth2Processing{})
	}

	c.setSameSiteCookie(ctx, accessTokenCookieName, token.AccessToken)
	c.setSameSiteCookie(ctx, refreshTokenCookieName, token.RefreshToken)

	return ctx.JSON(http.StatusOK, &oauth2Models.LoginResponse{Message: "Login Success"})
}

func (c *googleOAuth2Controller) Logout(ctx echo.Context) error {
	accesToken, err := ctx.Cookie(accessTokenCookieName)
	if err != nil {
		c.logger.Errorf("Failed to read acces token: %s", err.Error())
		return custom.Error(ctx, http.StatusBadRequest, &oauth2Exceptions.Logout{})
	}

	if err := c.revokeToken(accesToken.Value); err != nil {
		c.logger.Errorf("Failed to revoke access token: %s", err.Error())
		return custom.Error(ctx, http.StatusInternalServerError, &oauth2Exceptions.Logout{})
	}

	c.removeSameSiteCookie(ctx, accessTokenCookieName)
	c.removeSameSiteCookie(ctx, refreshTokenCookieName)

	return ctx.JSON(http.StatusOK, &oauth2Models.LogoutResponse{Message: "Logout Successful"})
}

func (c *googleOAuth2Controller) getUserInfo(client *http.Client) (*oauth2Models.UserInfo, error) {
	res, err := client.Get(c.oauth2Conf.UserInfoUrl)
	if err != nil {
		c.logger.Errorf("Failed to get user info: %s", err.Error())
		return nil, err
	}

	defer res.Body.Close()

	userInfoBytes, err := io.ReadAll(res.Body)
	if err != nil {
		c.logger.Errorf("Failed to read user info: %s", err.Error())
		return nil, err
	}

	userInfo := new(oauth2Models.UserInfo)
	if err := json.Unmarshal(userInfoBytes, &userInfo); err != nil {
		c.logger.Errorf("Failed to unmarshal user info: %s", err.Error())
		return nil, err
	}

	return userInfo, nil
}

func (c *googleOAuth2Controller) callbackValidate(ctx echo.Context) error {
	state := ctx.QueryParam("state")

	stateFromCookie, err := ctx.Cookie(stateCookieName)
	if err != nil {
		c.logger.Errorf("Failed to get state from cookie: %s", err.Error())
		return &oauth2Exceptions.Unauthorized{}
	}

	if state != stateFromCookie.Value {
		c.logger.Errorf("Invalid State: %s", state)
		return &oauth2Exceptions.Unauthorized{}
	}

	c.removeCookie(ctx, stateCookieName)

	return nil
}

func (c *googleOAuth2Controller) revokeToken(accessToken string) error {
	revokeURL := fmt.Sprintf("%s?token=%s", c.oauth2Conf.RevokeUrl, accessToken)

	res, err := http.Post(revokeURL, "application/x-www-from-urlencoded", nil)
	if err != nil {
		fmt.Println("Failed to revoke token: ", err)
		return err
	}

	defer res.Body.Close()

	return nil
}

func (c *googleOAuth2Controller) setCookie(ctx echo.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	}

	ctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) removeCookie(ctx echo.Context, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}

	ctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) setSameSiteCookie(ctx echo.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	ctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) removeSameSiteCookie(ctx echo.Context, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
	}

	ctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) randomState() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
