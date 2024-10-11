package services

import (
	itemShopModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
)

type ItemShopService interface {
	Listing(itemFilter *itemShopModels.ItemFilter) (*itemShopModels.ItemResult, error)
}
