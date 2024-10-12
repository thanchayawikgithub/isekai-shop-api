package repositories

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	itemShopModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
)

type ItemShopRepository interface {
	Listing(itemFilter *itemShopModels.ItemFilter) ([]*entities.Item, error)
	Counting(itemFilter *itemShopModels.ItemFilter) (int64, error)
	FindByID(itemID uint64) (*entities.Item, error)
}
