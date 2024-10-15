package models

import itemShopModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"

type (
	Inventory struct {
		Item     *itemShopModels.Item `json:"item"`
		Quantity uint                 `json:"quantity"`
	}

	ItemQuantityCounting struct {
		ItemID   uint64
		Quantity uint
	}
)
