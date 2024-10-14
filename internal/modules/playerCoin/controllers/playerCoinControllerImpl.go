package controllers

import (
	_ "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
	playerCoinServices "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/services"
)

type playerCoinControllerImpl struct {
	playerCoinService playerCoinServices.PlayerCoinService
}

func NewItemShopControllerImpl(playerCoinService playerCoinServices.PlayerCoinService) PlayerCoinController {
	return &playerCoinControllerImpl{playerCoinService}
}
