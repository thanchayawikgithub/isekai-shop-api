package repositories

import "github.com/thanchayawikgithub/isekai-shop-api/internal/entities"

type ItemShopRepository interface {
	Listing() ([]*entities.Item, error)
}
