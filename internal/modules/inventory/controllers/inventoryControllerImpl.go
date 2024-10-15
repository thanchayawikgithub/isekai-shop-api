package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	inventoryServices "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/inventory/services"
	"github.com/thanchayawikgithub/isekai-shop-api/pkg/custom"
	"github.com/thanchayawikgithub/isekai-shop-api/pkg/utils"
)

type inventoryControllerImpl struct {
	inventoryService inventoryServices.InventoryService
}

func NewInventoryControllerImpl(inventoryService inventoryServices.InventoryService) InventoryController {
	return &inventoryControllerImpl{inventoryService}
}

func (c *inventoryControllerImpl) Listing(ctx echo.Context) error {
	playerID, err := utils.GetReqPlayerID(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err)
	}

	inventoryListing, err := c.inventoryService.Listing(playerID)
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, inventoryListing)
}
