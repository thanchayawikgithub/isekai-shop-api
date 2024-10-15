package repositories

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	itemShopModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
	"gorm.io/gorm"
)

type ItemShopRepository interface {
	TransactionBegin() *gorm.DB
	TransactionRollback(tx *gorm.DB) error
	TransactionCommit(tx *gorm.DB) error
	Listing(itemFilter *itemShopModels.ItemFilter) ([]*entities.Item, error)
	Counting(itemFilter *itemShopModels.ItemFilter) (int64, error)
	FindByID(itemID uint64) (*entities.Item, error)
	FindByIDList(itemIDList []uint64) ([]*entities.Item, error)
	PurchaseHistoryRecording(tx *gorm.DB, purchaseEntity *entities.PurchaseHistory) (*entities.PurchaseHistory, error)
}
