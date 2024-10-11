package controllers

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/services"
)

type itemManagingControllerImpl struct {
	itemManagingService services.ItemManagingService
}

func NewItemManagingControllerImpl(itemManagingService services.ItemManagingService) ItemManagingController {
	return &itemManagingControllerImpl{itemManagingService}
}
