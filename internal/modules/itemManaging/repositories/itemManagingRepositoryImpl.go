package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	itemManagingException "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/exceptions"
	itemManagingModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/models"
)

type itemManagingRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db databases.Database, logger echo.Logger) ItemManagingRepository {
	return &itemManagingRepositoryImpl{db, logger}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	savedItem := new(entities.Item)
	if err := r.db.Connect().Create(itemEntity).Scan(savedItem).Error; err != nil {
		r.logger.Errorf("Failed to creating item: %s", err.Error())
		return nil, &itemManagingException.ItemCreating{}
	}
	return savedItem, nil
}

func (r *itemManagingRepositoryImpl) Editing(itemID uint64, itemEditingReq *itemManagingModels.ItemEditingReq) (uint64, error) {
	if err := r.db.Connect().Model(&entities.Item{}).Where("id = ?", itemID).Updates(itemEditingReq).Error; err != nil {
		r.logger.Errorf("Failed to editing item", err.Error())
		return 0, &itemManagingException.ItemEditing{}
	}

	return itemID, nil
}

func (r *itemManagingRepositoryImpl) Archiving(itemID uint64) error {
	if err := r.db.Connect().Table("items").Where("id = ?", itemID).Update("is_archive", true).Error; err != nil {
		r.logger.Errorf("Archiving item failed: %s", err.Error())
		return &itemManagingException.ItemArchiving{ItemID: itemID}
	}

	return nil
}
