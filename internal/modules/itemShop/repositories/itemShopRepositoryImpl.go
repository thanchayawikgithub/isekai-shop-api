package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/exceptions"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
	"gorm.io/gorm"
)

type itemShopRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *models.ItemFilter) ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	query := r.db.Model(&entities.Item{})

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	if result := query.Find(&itemList); result.Error != nil {
		r.logger.Errorf("Failed to list items: %s", result.Error)
		return nil, &exceptions.ItemListing{}
	}

	return itemList, nil
}
