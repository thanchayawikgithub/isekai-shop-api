package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	playerCoinModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/models"
	playerCoinServices "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/services"
	"github.com/thanchayawikgithub/isekai-shop-api/pkg/custom"
	"github.com/thanchayawikgithub/isekai-shop-api/pkg/utils"
)

type playerCoinControllerImpl struct {
	playerCoinService playerCoinServices.PlayerCoinService
}

func NewPlayerCoinControllerImpl(playerCoinService playerCoinServices.PlayerCoinService) PlayerCoinController {
	return &playerCoinControllerImpl{playerCoinService}
}

func (c *playerCoinControllerImpl) CoinAdding(ctx echo.Context) error {
	playerID, err := utils.GetReqPlayerID(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err)
	}

	coinAddingReq := new(playerCoinModels.CoinAddingReq)

	customRequest := custom.NewCustomRequest(ctx)
	if err := customRequest.Bind(coinAddingReq); err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err)
	}
	coinAddingReq.PlayerID = playerID

	savedPlayerCoin, err := c.playerCoinService.CoinAdding(coinAddingReq)
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, savedPlayerCoin)
}

func (c *playerCoinControllerImpl) CoinShowing(ctx echo.Context) error {
	playerID, err := utils.GetReqPlayerID(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err)
	}

	playerCoinShowing := c.playerCoinService.CoinShowing(playerID)

	return ctx.JSON(http.StatusOK, playerCoinShowing)
}
