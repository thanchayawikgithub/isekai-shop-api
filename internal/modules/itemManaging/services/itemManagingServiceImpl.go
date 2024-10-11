package services

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/repositories"
)

type itemManagingServiceImpl struct {
	itemManagingRepo repositories.ItemManagingRepository
}

func NewItemManagingServiceImpl(itemManagingRepo repositories.ItemManagingRepository) ItemManagingService {
	return &itemManagingServiceImpl{itemManagingRepo}
}
