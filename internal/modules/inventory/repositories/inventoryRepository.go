package repositories

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
)

type InventoryRepository interface {
	Filling(inventoryEntities []*entities.Inventory) ([]*entities.Inventory, error)
	Removing(playerID string, itemID uint64, quantity int) error
	PlayerItemCounting(playerID string, itemID uint64) int64
	Listing(playerID string) ([]*entities.Inventory, error)
}
