package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	itemShopExceptions "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/exceptions"
	itemShopModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
	"gorm.io/gorm"
)

type itemShopRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *itemShopModels.ItemFilter) ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false)

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	limit := int(itemFilter.Size)

	if result := query.Offset(offset).Limit(limit).Order("id desc").Find(&itemList); result.Error != nil {
		r.logger.Errorf("Failed to list items: %s", result.Error)
		return nil, &itemShopExceptions.ItemListing{}
	}

	return itemList, nil
}

func (r *itemShopRepositoryImpl) Counting(itemFilter *itemShopModels.ItemFilter) (int64, error) {
	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false)

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	var count int64

	if result := query.Count(&count); result.Error != nil {
		r.logger.Errorf("Failed to counting items: %s", result.Error)
		return -1, &itemShopExceptions.ItemCounting{}
	}

	return count, nil
}
