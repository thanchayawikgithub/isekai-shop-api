package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	itemManagingException "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/exceptions"
	"gorm.io/gorm"
)

type itemManagingRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemManagingRepository {
	return &itemManagingRepositoryImpl{db, logger}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	savedItem := new(entities.Item)
	if err := r.db.Create(itemEntity).Scan(savedItem).Error; err != nil {
		r.logger.Errorf("Failed to creating item: %s", err.Error())
		return nil, &itemManagingException.ItemCreating{}
	}
	return savedItem, nil
}
