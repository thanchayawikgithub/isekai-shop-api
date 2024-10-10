package services

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/repositories"
)

type itemShopServiceImpl struct {
	itemShopRepo repositories.ItemShopRepository
}

func NewItemShopServiceImpl(itemShopRepo repositories.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepo}
}

func (s *itemShopServiceImpl) Listing() ([]*models.Item, error) {
	itemList, err := s.itemShopRepo.Listing()

	if err != nil {
		return nil, err
	}

	itemModelList := make([]*models.Item, 0)
	for _, item := range itemList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return itemModelList, nil
}
