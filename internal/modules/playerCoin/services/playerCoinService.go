package services

import (
	playerCoinModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/models"
)

type PlayerCoinService interface {
	CoinAdding(coinAddingReq *playerCoinModels.CoinAddingReq) (*playerCoinModels.PlayerCoin, error)
	CoinShowing(playerID string) *playerCoinModels.PlayerCoinShowing
}
