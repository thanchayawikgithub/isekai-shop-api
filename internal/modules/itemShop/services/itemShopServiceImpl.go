package services

import "github.com/thanchayawikgithub/isekai-shop-api/internal/itemShop/repositories"

type itemShopServiceImpl struct {
	itemShopRepo repositories.ItemShopRepository
}

func NewItemShopServiceImpl(itemShopRepo repositories.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepo}
}
