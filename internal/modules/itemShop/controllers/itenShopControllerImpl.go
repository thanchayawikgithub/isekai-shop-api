package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/services"
)

type itemShopControllerImpl struct {
	itemShopService services.ItemShopService
}

func NewItemShopControllerImpl(itemShopService services.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(ctx echo.Context) error {
	itemModelList, err := c.itemShopService.Listing()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, itemModelList)
}
