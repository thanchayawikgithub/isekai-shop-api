package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	inventoryExceptions "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/inventory/exceptions"
)

type inventoryRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db databases.Database, logger echo.Logger) InventoryRepository {
	return &inventoryRepositoryImpl{db, logger}
}

func (r *inventoryRepositoryImpl) Filling(inventoryEntities []*entities.Inventory) ([]*entities.Inventory, error) {
	savedInventories := make([]*entities.Inventory, 0)

	if err := r.db.Connect().CreateInBatches(inventoryEntities, len(inventoryEntities)).Scan(&savedInventories).Error; err != nil {
		r.logger.Errorf("Error filling inventory: %v", err)
		return nil, &inventoryExceptions.InventoryFilling{
			PlayerID: inventoryEntities[0].PlayerID,
			ItemID:   inventoryEntities[0].ItemID,
		}
	}

	return savedInventories, nil
}

func (r *inventoryRepositoryImpl) Removing(playerID string, itemID uint64, quantity int) error {
	inventories, err := r.findPlayerItem(playerID, itemID, quantity)
	if err != nil {
		return err
	}

	tx := r.db.Connect().Begin()
	for _, inventory := range inventories {
		inventory.IsDeleted = true

		if err := tx.Model(&entities.Inventory{}).Where("id = ?", inventory.ID).Updates(inventory); err != nil {
			tx.Rollback()
			r.logger.Errorf("error removing player item in inventory: %s", err)
			return &inventoryExceptions.PlayerItemRemoving{ItemID: itemID}
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		r.logger.Errorf("error removing player item in inventory: %s", err.Error())
		return &inventoryExceptions.PlayerItemRemoving{ItemID: itemID}
	}

	return nil
}

func (r *inventoryRepositoryImpl) PlayerItemCounting(playerID string, itemID uint64) int64 {
	var count int64

	if err := r.db.Connect().Model(&entities.Inventory{}).
		Where("player_id = ? and item_id = ? and is_deleted = ?", playerID, itemID, false).
		Count(&count).Error; err != nil {
		r.logger.Errorf("error counting player item in inventory: %s", err.Error())
		return -1
	}

	return count
}

func (r *inventoryRepositoryImpl) findPlayerItem(playerID string, itemID uint64, quantity int) ([]*entities.Inventory, error) {
	inventories := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where("player_id = ? and item_id = ? and is_deleted = ?", playerID, itemID, false).
		Limit(quantity).Find(inventories).Error; err != nil {
		r.logger.Errorf("error finding player item in inventory by ID: %s", err.Error())
		return nil, &inventoryExceptions.PlayerItemRemoving{ItemID: itemID}
	}

	return inventories, nil
}

func (r *inventoryRepositoryImpl) Listing(playerID string) ([]*entities.Inventory, error) {
	inventories := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where("player_id = ? and is_deleted = ?", playerID, false).Find(&inventories).Error; err != nil {
		r.logger.Errorf("error listing player inventory: %s", err.Error())
		return nil, &inventoryExceptions.PlayerItemsFinding{PlayerID: playerID}
	}

	return inventories, nil
}
