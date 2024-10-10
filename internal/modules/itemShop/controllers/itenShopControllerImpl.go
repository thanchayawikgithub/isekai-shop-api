package controllers

import "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/services"

type itemShopControllerImpl struct {
	itemShopService services.ItemShopService
}

func NewItemShopControllerImpl(itemShopService services.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}
