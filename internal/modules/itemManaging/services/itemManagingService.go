package services

import (
	itemManagingModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/models"
	itemShopModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
)

type ItemManagingService interface {
	Creating(itemCreatingReq *itemManagingModels.ItemCreatingReq) (*itemShopModels.Item, error)
	Editing(itemID uint64, itemEditingReq *itemManagingModels.ItemEditingReq) (*itemShopModels.Item, error)
	Archiving(itemID uint64) error
}
