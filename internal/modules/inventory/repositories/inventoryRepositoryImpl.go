package repositories

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	inventoryExceptions "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/inventory/exceptions"
	"gorm.io/gorm"
)

type inventoryRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db databases.Database, logger echo.Logger) InventoryRepository {
	return &inventoryRepositoryImpl{db, logger}
}

func (r *inventoryRepositoryImpl) Filling(tx *gorm.DB, playerID string, itemID uint64, quantity int) ([]*entities.Inventory, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	inventoryEntities := make([]*entities.Inventory, 0)

	for range quantity {
		inventoryEntities = append(inventoryEntities, &entities.Inventory{
			PlayerID: playerID,
			ItemID:   itemID,
		})
	}

	if err := conn.Create(inventoryEntities).Error; err != nil {
		r.logger.Errorf("Error filling inventory: %s", err.Error())
		return nil, &inventoryExceptions.InventoryFilling{
			PlayerID: playerID,
			ItemID:   itemID,
		}
	}

	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) Removing(tx *gorm.DB, playerID string, itemID uint64, quantity int) error {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	inventories, err := r.findPlayerItem(playerID, itemID, quantity)
	if err != nil {
		return err
	}

	for _, inventory := range inventories {
		inventory.IsDeleted = true

		fmt.Print(*inventory)

		if err := conn.Model(&entities.Inventory{}).Where("id = ?", inventory.ID).Updates(inventory).Error; err != nil {
			r.logger.Errorf("error removing player item in inventory: %s", err)
			return &inventoryExceptions.PlayerItemRemoving{ItemID: itemID}
		}
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
		Limit(quantity).Find(&inventories).Error; err != nil {
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
