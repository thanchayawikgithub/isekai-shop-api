package services

import (
	_ "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
	playerCoinRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/repositories"
)

type playerCoinServiceImpl struct {
	playerCoinRepo playerCoinRepositories.PlayerCoinRepository
}

func NewItemShopServiceImpl(playerCoinRepo playerCoinRepositories.PlayerCoinRepository) PlayerCoinService {
	return &playerCoinServiceImpl{playerCoinRepo}
}
