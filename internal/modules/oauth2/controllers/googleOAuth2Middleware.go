package controllers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	oauth2Exceptions "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/oauth2/exceptions"
	"github.com/thanchayawikgithub/isekai-shop-api/pkg/custom"
	"golang.org/x/oauth2"
)

func (c *googleOAuth2Controller) PlayerAuthorize(ctx echo.Context, next echo.HandlerFunc) error {
	context := context.Background()
	tokenSource, err := c.getTokenSource(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusUnauthorized, err)
	}

	if !tokenSource.Valid() {
		tokenSource, err = c.refreshPlayerToken(ctx, tokenSource)
		if err != nil {
			return custom.Error(ctx, http.StatusUnauthorized, err)
		}
	}

	client := playerGoogleOAuth2.Client(context, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		return custom.Error(ctx, http.StatusUnauthorized, err)
	}

	if !c.oauth2Service.IsPlayer(userInfo.ID) {
		return custom.Error(ctx, http.StatusUnauthorized, &oauth2Exceptions.Unauthorized{})
	}

	ctx.Set("playerID", userInfo.ID)

	return next(ctx)
}

func (c *googleOAuth2Controller) AdminAuthorize(ctx echo.Context, next echo.HandlerFunc) error {
	context := context.Background()
	tokenSource, err := c.getTokenSource(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusUnauthorized, err)
	}

	if !tokenSource.Valid() {
		tokenSource, err = c.refreshPlayerToken(ctx, tokenSource)
		if err != nil {
			return custom.Error(ctx, http.StatusUnauthorized, err)
		}
	}

	client := adminGoogleOAuth2.Client(context, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		return custom.Error(ctx, http.StatusUnauthorized, err)
	}

	if !c.oauth2Service.IsAdmin(userInfo.ID) {
		return custom.Error(ctx, http.StatusUnauthorized, &oauth2Exceptions.Unauthorized{})
	}

	ctx.Set("adminID", userInfo.ID)

	return next(ctx)
}

func (c *googleOAuth2Controller) refreshPlayerToken(ctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	context := context.Background()

	updatedToken, err := playerGoogleOAuth2.TokenSource(context, token).Token()
	if err != nil {
		return nil, &oauth2Exceptions.Unauthorized{}
	}

	c.setSameSiteCookie(ctx, accessTokenCookieName, updatedToken.AccessToken)
	c.setSameSiteCookie(ctx, refreshTokenCookieName, updatedToken.RefreshToken)

	return updatedToken, nil
}

func (c *googleOAuth2Controller) refreshAdminToken(ctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	context := context.Background()

	updatedToken, err := adminGoogleOAuth2.TokenSource(context, token).Token()
	if err != nil {
		return nil, &oauth2Exceptions.Unauthorized{}
	}

	c.setSameSiteCookie(ctx, accessTokenCookieName, updatedToken.AccessToken)
	c.setSameSiteCookie(ctx, refreshTokenCookieName, updatedToken.RefreshToken)

	return updatedToken, nil
}

func (c *googleOAuth2Controller) getTokenSource(ctx echo.Context) (*oauth2.Token, error) {
	accesToken, err := ctx.Cookie(accessTokenCookieName)
	if err != nil {
		return nil, &oauth2Exceptions.Unauthorized{}
	}

	refreshToken, err := ctx.Cookie(refreshTokenCookieName)
	if err != nil {
		return nil, &oauth2Exceptions.Unauthorized{}
	}

	return &oauth2.Token{AccessToken: accesToken.Value, RefreshToken: refreshToken.Value}, nil
}
