package repositories

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	Filling(tx *gorm.DB, playerID string, itemID uint64, quantity int) ([]*entities.Inventory, error)
	Removing(tx *gorm.DB, playerID string, itemID uint64, quantity int) error
	PlayerItemCounting(playerID string, itemID uint64) int64
	Listing(playerID string) ([]*entities.Inventory, error)
}
