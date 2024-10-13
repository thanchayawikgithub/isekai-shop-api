package controllers

import (
	"net/http"
	"strconv"

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
		return custom.Error(ctx, http.StatusBadRequest, err)
	}

	savedItem, err := c.itemManagingService.Creating(itemCreatingReq)
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, savedItem)
}

func (c *itemManagingControllerImpl) Editing(ctx echo.Context) error {
	itemID, err := c.getItemID(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err)
	}

	itemEditingReq := new(itemManagingModels.ItemEditingReq)

	customRequest := custom.NewCustomRequest(ctx)

	if err := customRequest.Bind(itemEditingReq); err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err)
	}

	item, err := c.itemManagingService.Editing(itemID, itemEditingReq)
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, item)
}

func (c *itemManagingControllerImpl) Archiving(ctx echo.Context) error {
	itemID, err := c.getItemID(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err)
	}

	if err := c.itemManagingService.Archiving(itemID); err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *itemManagingControllerImpl) getItemID(ctx echo.Context) (uint64, error) {
	itemIDStr := ctx.Param("itemID")
	itemID, err := strconv.ParseUint(itemIDStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return itemID, nil
}
