package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
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

	if err := r.db.Find(&itemList).Error; err != nil {
		r.logger.Errorf("Failed to list items: %s", err.Error())
		return nil, err
	}

	return itemList, nil
}
