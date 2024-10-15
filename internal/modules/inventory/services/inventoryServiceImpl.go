package services

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	inventoryModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/inventory/models"
	inventoryRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/inventory/repositories"
	itemShopRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/repositories"
)

type inventoryServiceImpl struct {
	inventoryRepo inventoryRepositories.InventoryRepository
	itemShopRepo  itemShopRepositories.ItemShopRepository
}

func NewInventoryServiceImpl(inventoryRepo inventoryRepositories.InventoryRepository, itemShopRepo itemShopRepositories.ItemShopRepository) InventoryService {
	return &inventoryServiceImpl{inventoryRepo, itemShopRepo}
}

func (s *inventoryServiceImpl) Listing(playerID string) ([]*inventoryModels.Inventory, error) {
	inventories, err := s.inventoryRepo.Listing(playerID)
	if err != nil {
		return nil, err
	}

	itemQuantityCounting := s.getUniqueItemsWithQuantity(inventories)

	return s.buildInventoryListingResult(itemQuantityCounting), nil
}

func (s *inventoryServiceImpl) getUniqueItemsWithQuantity(inventories []*entities.Inventory) []inventoryModels.ItemQuantityCounting {
	itemQuantityCounting := make([]inventoryModels.ItemQuantityCounting, 0)
	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range inventories {
		itemMapWithQuantity[inventory.ItemID]++
	}

	for itemID, quantity := range itemMapWithQuantity {
		itemQuantityCounting = append(itemQuantityCounting, inventoryModels.ItemQuantityCounting{ItemID: itemID, Quantity: quantity})
	}

	return itemQuantityCounting
}

func (s *inventoryServiceImpl) buildInventoryListingResult(itemQuantityCounting []inventoryModels.ItemQuantityCounting) []*inventoryModels.Inventory {
	uniqueItemIDList := s.getItemID(itemQuantityCounting)

	items, err := s.itemShopRepo.FindByIDList(uniqueItemIDList)
	if err != nil {
		return make([]*inventoryModels.Inventory, 0)
	}

	results := make([]*inventoryModels.Inventory, 0)
	itemMapWithQuantity := s.getItemMapWithQuantity(itemQuantityCounting)

	for _, item := range items {
		results = append(results, &inventoryModels.Inventory{Item: item.ToItemModel(), Quantity: itemMapWithQuantity[item.ID]})
	}

	return results
}

func (s *inventoryServiceImpl) getItemID(itemQuantityCounting []inventoryModels.ItemQuantityCounting) []uint64 {
	uniqueItemIDList := make([]uint64, 0)

	for _, inventory := range itemQuantityCounting {
		uniqueItemIDList = append(uniqueItemIDList, inventory.ItemID)
	}

	return uniqueItemIDList
}

func (s *inventoryServiceImpl) getItemMapWithQuantity(itemQuantityCounting []inventoryModels.ItemQuantityCounting) map[uint64]uint {
	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range itemQuantityCounting {
		itemMapWithQuantity[inventory.ItemID] = inventory.Quantity
	}

	return itemMapWithQuantity
}
