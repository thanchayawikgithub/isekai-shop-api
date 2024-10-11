package repositories

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
)

type ItemShopRepository interface {
	Listing(itemFilter *models.ItemFilter) ([]*entities.Item, error)
}
