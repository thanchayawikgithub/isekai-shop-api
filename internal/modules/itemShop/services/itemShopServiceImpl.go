package services

import (
	"math"

	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	inventoryRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/inventory/repositories"
	itemShopExceptions "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/exceptions"
	itemShopModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
	itemShopRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/repositories"
	playerCoinModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/models"
	playerCoinRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/repositories"
)

type itemShopServiceImpl struct {
	itemShopRepo   itemShopRepositories.ItemShopRepository
	playerCoinRepo playerCoinRepositories.PlayerCoinRepository
	inventoryRepo  inventoryRepositories.InventoryRepository
}

func NewItemShopServiceImpl(itemShopRepo itemShopRepositories.ItemShopRepository,
	playerCoinRepo playerCoinRepositories.PlayerCoinRepository,
	inventoryRepo inventoryRepositories.InventoryRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepo, playerCoinRepo, inventoryRepo}
}

func (s *itemShopServiceImpl) Listing(itemFilter *itemShopModels.ItemFilter) (*itemShopModels.ItemResult, error) {
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

func (s *itemShopServiceImpl) Buying(buyingReq *itemShopModels.BuyingReq) (*playerCoinModels.PlayerCoin, error) {
	itemEntity, err := s.itemShopRepo.FindByID(buyingReq.ItemID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.totalPriceCalculation(itemEntity.ToItemModel(), buyingReq.Quantity)

	if err := s.playerCoinChecking(buyingReq.PlayerID, totalPrice); err != nil {
		return nil, err
	}

	tx := s.itemShopRepo.TransactionBegin()

	_, err = s.itemShopRepo.PurchaseHistoryRecording(tx, &entities.PurchaseHistory{
		PlayerID:        buyingReq.PlayerID,
		ItemID:          buyingReq.ItemID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		Quantity:        buyingReq.Quantity,
		IsBuying:        true,
	})

	if err != nil {
		s.itemShopRepo.TransactionRollback(tx)
		return nil, err
	}

	playerCoin, err := s.playerCoinRepo.CoinAdding(tx, &entities.PlayerCoin{
		PlayerID: buyingReq.PlayerID,
		Amount:   -totalPrice,
	})
	if err != nil {
		s.itemShopRepo.TransactionRollback(tx)
	}

	_, err = s.inventoryRepo.Filling(tx, buyingReq.PlayerID, buyingReq.ItemID, int(buyingReq.Quantity))
	if err != nil {
		s.itemShopRepo.TransactionRollback(tx)
		return nil, err
	}

	if err := s.itemShopRepo.TransactionCommit(tx); err != nil {
		return nil, err
	}

	return playerCoin.ToPlayerCoinModel(), nil
}

func (s *itemShopServiceImpl) Selling(sellingReq *itemShopModels.SellingReq) (*playerCoinModels.PlayerCoin, error) {
	itemEntity, err := s.itemShopRepo.FindByID(sellingReq.ItemID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.totalPriceCalculation(itemEntity.ToItemModel(), sellingReq.Quantity)
	totalPrice = totalPrice / 2

	if err := s.playerItemChecking(sellingReq.PlayerID, sellingReq.ItemID, sellingReq.Quantity); err != nil {
		return nil, err
	}

	tx := s.itemShopRepo.TransactionBegin()

	_, err = s.itemShopRepo.PurchaseHistoryRecording(tx, &entities.PurchaseHistory{
		PlayerID:        sellingReq.PlayerID,
		ItemID:          sellingReq.ItemID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		Quantity:        sellingReq.Quantity,
		IsBuying:        false,
	})
	if err != nil {
		s.itemShopRepo.TransactionRollback(tx)
		return nil, err
	}

	playerCoin, err := s.playerCoinRepo.CoinAdding(tx, &entities.PlayerCoin{
		PlayerID: sellingReq.PlayerID,
		Amount:   totalPrice,
	})
	if err != nil {
		s.itemShopRepo.TransactionRollback(tx)
	}

	if err = s.inventoryRepo.Removing(tx, sellingReq.PlayerID, sellingReq.ItemID, int(sellingReq.Quantity)); err != nil {
		s.itemShopRepo.TransactionRollback(tx)
		return nil, err
	}

	if err := s.itemShopRepo.TransactionCommit(tx); err != nil {
		return nil, err
	}

	return playerCoin.ToPlayerCoinModel(), nil
}

func (s *itemShopServiceImpl) calTotalPage(totalItems int64, size int64) int64 {
	return int64(math.Ceil(float64(totalItems) / float64(size)))
}

func (s *itemShopServiceImpl) toItemResult(itemList []*entities.Item, page, totalPage int64) *itemShopModels.ItemResult {
	itemModelList := make([]*itemShopModels.Item, 0)
	for _, item := range itemList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return &itemShopModels.ItemResult{
		Items: itemModelList,
		Paginate: itemShopModels.PaginateResult{
			Page:     page,
			TotaPage: totalPage,
		},
	}
}

func (s *itemShopServiceImpl) totalPriceCalculation(item *itemShopModels.Item, quantity uint) int64 {
	return int64(item.Price) * int64(quantity)
}

func (s *itemShopServiceImpl) playerCoinChecking(playerID string, totalPrice int64) error {
	playerCoin, err := s.playerCoinRepo.CoinShowing(playerID)
	if err != nil {
		return err
	}

	if playerCoin.Coin < totalPrice {
		return &itemShopExceptions.CoinNotEnough{}
	}

	return nil
}

func (s *itemShopServiceImpl) playerItemChecking(playerID string, itemID uint64, quantity uint) error {
	itemCounting := s.inventoryRepo.PlayerItemCounting(playerID, itemID)

	if int(itemCounting) < int(quantity) {
		return &itemShopExceptions.ItemNotEnough{ItemID: itemID}
	}

	return nil
}
