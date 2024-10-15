package services

import (
	inventoryModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/inventory/models"
)

type InventoryService interface {
	Listing(playerID string) ([]*inventoryModels.Inventory, error)
}
