package services

import "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"

type ItemShopService interface {
	Listing() ([]*models.Item, error)
}
