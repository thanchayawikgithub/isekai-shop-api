package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	itemShopExceptions "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/exceptions"
	itemShopModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
	"gorm.io/gorm"
)

type itemShopRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db databases.Database, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) TransactionBegin() *gorm.DB {
	tx := r.db.Connect()
	return tx.Begin()
}

func (r *itemShopRepositoryImpl) TransactionRollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *itemShopRepositoryImpl) TransactionCommit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *itemShopModels.ItemFilter) ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	query := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ?", false)

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	limit := int(itemFilter.Size)

	if result := query.Offset(offset).Limit(limit).Order("id desc").Find(&itemList); result.Error != nil {
		r.logger.Errorf("Failed to list items: %s", result.Error)
		return nil, &itemShopExceptions.ItemListing{}
	}

	return itemList, nil
}

func (r *itemShopRepositoryImpl) Counting(itemFilter *itemShopModels.ItemFilter) (int64, error) {
	query := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ?", false)

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	var count int64

	if result := query.Count(&count); result.Error != nil {
		r.logger.Errorf("Failed to counting items: %s", result.Error)
		return -1, &itemShopExceptions.ItemCounting{}
	}

	return count, nil
}

func (r *itemShopRepositoryImpl) FindByID(itemID uint64) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.Connect().First(item, itemID).Error; err != nil {
		r.logger.Errorf("Failed to find item by ID: %s", err.Error())
		return nil, &itemShopExceptions.ItemNotFound{}
	}

	return item, nil
}

func (r *itemShopRepositoryImpl) FindByIDList(itemIDList []uint64) ([]*entities.Item, error) {
	items := make([]*entities.Item, 0)

	if err := r.db.Connect().Model(&entities.Item{}).Where("id in ?", itemIDList).Find(&items).Error; err != nil {
		r.logger.Errorf("Failed to find item by ID list: %s", err.Error())
		return nil, &itemShopExceptions.ItemListing{}
	}

	return items, nil
}

func (r *itemShopRepositoryImpl) PurchaseHistoryRecording(tx *gorm.DB, purchaseEntity *entities.PurchaseHistory) (*entities.PurchaseHistory, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	savedPurchase := new(entities.PurchaseHistory)

	if err := conn.Create(purchaseEntity).Scan(savedPurchase).Error; err != nil {
		r.logger.Errorf("Failed to record purchase history: %s", err.Error())
		return nil, &itemShopExceptions.PurchaseHistoryRecording{}
	}

	return savedPurchase, nil
}
