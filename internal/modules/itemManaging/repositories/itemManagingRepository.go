package repositories

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	itemManagingModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/models"
)

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
	Editing(itemID uint64, itemEditingReq *itemManagingModels.ItemEditingReq) (uint64, error)
	Archiving(itemID uint64) error
}
