package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	itemManagingModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/models"
	itemManagingServices "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/services"
	"github.com/thanchayawikgithub/isekai-shop-api/pkg/custom"
)

type itemManagingControllerImpl struct {
	itemManagingService itemManagingServices.ItemManagingService
}

func NewItemManagingControllerImpl(itemManagingService itemManagingServices.ItemManagingService) ItemManagingController {
	return &itemManagingControllerImpl{itemManagingService}
}

func (c *itemManagingControllerImpl) Creating(ctx echo.Context) error {
	itemCreatingReq := new(itemManagingModels.ItemCreatingReq)

	customRequest := custom.NewCustomRequest(ctx)
	if err := customRequest.Bind(itemCreatingReq); err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	savedItem, err := c.itemManagingService.Creating(itemCreatingReq)
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, savedItem)
}
