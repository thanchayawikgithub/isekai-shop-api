package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	itemShopModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
	itemShopServices "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/services"
	"github.com/thanchayawikgithub/isekai-shop-api/pkg/custom"
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
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	itemModelList, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, itemModelList)

}
