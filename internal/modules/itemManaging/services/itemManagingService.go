package services

import (
	itemManagingModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/models"
	itemShopModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
)

type ItemManagingService interface {
	Creating(itemCreatingReq *itemManagingModels.ItemCreatingReq) (*itemShopModels.Item, error)
}
