package services

import (
	itemShopModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
	playerCoinModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/models"
)

type ItemShopService interface {
	Listing(itemFilter *itemShopModels.ItemFilter) (*itemShopModels.ItemResult, error)
	Buying(buyingReq *itemShopModels.BuyingReq) (*playerCoinModels.PlayerCoin, error)
	Selling(sellingReq *itemShopModels.SellingReq) (*playerCoinModels.PlayerCoin, error)
}
