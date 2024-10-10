package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/exceptions"
	"gorm.io/gorm"
)

type itemShopRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing() ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	if result := r.db.Find(&itemList); result.Error != nil {
		r.logger.Errorf("Failed to list items: %s", result.Error)
		return nil, &exceptions.ItemListing{}
	}

	return itemList, nil
}
