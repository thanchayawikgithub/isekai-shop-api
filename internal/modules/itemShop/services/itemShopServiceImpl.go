package services

import (
	"math"

	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/repositories"
)

type itemShopServiceImpl struct {
	itemShopRepo repositories.ItemShopRepository
}

func NewItemShopServiceImpl(itemShopRepo repositories.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepo}
}

func (s *itemShopServiceImpl) Listing(itemFilter *models.ItemFilter) (*models.ItemResult, error) {
	itemList, err := s.itemShopRepo.Listing(itemFilter)
	if err != nil {
		return nil, err
	}

	itemCount, err := s.itemShopRepo.Counting(itemFilter)
	if err != nil {
		return nil, err
	}

	totalPage := s.calTotalPage(itemCount, itemFilter.Size)

	return s.toItemResult(itemList, itemFilter.Page, totalPage), nil
}

func (s *itemShopServiceImpl) calTotalPage(totalItems int64, size int64) int64 {
	return int64(math.Ceil(float64(totalItems) / float64(size)))
}

func (s *itemShopServiceImpl) toItemResult(itemList []*entities.Item, page, totalPage int64) *models.ItemResult {
	itemModelList := make([]*models.Item, 0)
	for _, item := range itemList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return &models.ItemResult{
		Items: itemModelList,
		Paginate: models.PaginateResult{
			Page:     page,
			TotaPage: totalPage,
		},
	}
}
