package services

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"

	itemManagingModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/models"
	itemManagingRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/repositories"
	itemShopModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
	itemShopRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/repositories"
)

type itemManagingServiceImpl struct {
	itemManagingRepo itemManagingRepositories.ItemManagingRepository
	itemShopRepo     itemShopRepositories.ItemShopRepository
}

func NewItemManagingServiceImpl(
	itemManagingRepo itemManagingRepositories.ItemManagingRepository,
	itemShopRepo itemShopRepositories.ItemShopRepository,
) ItemManagingService {
	return &itemManagingServiceImpl{itemManagingRepo, itemShopRepo}
}

func (s *itemManagingServiceImpl) Creating(itemCreatingReq *itemManagingModels.ItemCreatingReq) (*itemShopModels.Item, error) {
	item := &entities.Item{
		Name:        itemCreatingReq.Name,
		Description: itemCreatingReq.Description,
		Picture:     itemCreatingReq.Picture,
		Price:       itemCreatingReq.Price,
	}

	savedItem, err := s.itemManagingRepo.Creating(item)
	if err != nil {
		return nil, err
	}

	return savedItem.ToItemModel(), nil
}

func (s *itemManagingServiceImpl) Editing(itemID uint64, itemEditingReq *itemManagingModels.ItemEditingReq) (*itemShopModels.Item, error) {
	_, err := s.itemManagingRepo.Editing(itemID, itemEditingReq)
	if err != nil {
		return nil, err
	}

	itemEntityResult, err := s.itemShopRepo.FindByID(itemID)
	if err != nil {
		return nil, err
	}

	return itemEntityResult.ToItemModel(), nil
}

func (s *itemManagingServiceImpl) Archiving(itemID uint64) error {
	return s.itemManagingRepo.Archiving(itemID)
}
