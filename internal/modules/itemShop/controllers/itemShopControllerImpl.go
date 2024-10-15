package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	itemShopModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
	itemShopServices "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/services"
	"github.com/thanchayawikgithub/isekai-shop-api/pkg/custom"
	"github.com/thanchayawikgithub/isekai-shop-api/pkg/utils"
)

type itemShopControllerImpl struct {
	itemShopService itemShopServices.ItemShopService
}

func NewItemShopControllerImpl(itemShopService itemShopServices.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(ctx echo.Context) error {
	itemFilter := new(itemShopModels.ItemFilter)

	customRequest := custom.NewCustomRequest(ctx)
	if err := customRequest.Bind(itemFilter); err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err)
	}

	itemModelList, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, itemModelList)
}

func (c *itemShopControllerImpl) Buying(ctx echo.Context) error {
	playerID, err := utils.GetReqPlayerID(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err)
	}

	buyingReq := new(itemShopModels.BuyingReq)

	customRequest := custom.NewCustomRequest(ctx)
	if err := customRequest.Bind(buyingReq); err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err)
	}
	buyingReq.PlayerID = playerID

	playerCoin, err := c.itemShopService.Buying(buyingReq)
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, playerCoin)
}

func (c *itemShopControllerImpl) Selling(ctx echo.Context) error {
	playerID, err := utils.GetReqPlayerID(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err)
	}

	sellingReq := new(itemShopModels.SellingReq)

	customRequest := custom.NewCustomRequest(ctx)
	if err := customRequest.Bind(sellingReq); err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err)
	}
	sellingReq.PlayerID = playerID

	playerCoin, err := c.itemShopService.Selling(sellingReq)
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, playerCoin)
}
