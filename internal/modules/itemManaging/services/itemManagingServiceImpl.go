package services

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	itemManagingModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/models"
	itemManagingRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/repositories"
	itemShopModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
)

type itemManagingServiceImpl struct {
	itemManagingRepo itemManagingRepositories.ItemManagingRepository
}

func NewItemManagingServiceImpl(itemManagingRepo itemManagingRepositories.ItemManagingRepository) ItemManagingService {
	return &itemManagingServiceImpl{itemManagingRepo}
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
