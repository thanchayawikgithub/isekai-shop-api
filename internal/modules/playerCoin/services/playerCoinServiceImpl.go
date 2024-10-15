package services

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	playerCoinModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/models"
	playerCoinRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/repositories"
)

type playerCoinServiceImpl struct {
	playerCoinRepo playerCoinRepositories.PlayerCoinRepository
}

func NewPlayerCoinServiceImpl(playerCoinRepo playerCoinRepositories.PlayerCoinRepository) PlayerCoinService {
	return &playerCoinServiceImpl{playerCoinRepo}
}

func (s *playerCoinServiceImpl) CoinAdding(coinAddingReq *playerCoinModels.CoinAddingReq) (*playerCoinModels.PlayerCoin, error) {
	playerCoinEntity := &entities.PlayerCoin{
		PlayerID: coinAddingReq.PlayerID,
		Amount:   coinAddingReq.Amount,
	}

	savedPlayerCoin, err := s.playerCoinRepo.CoinAdding(playerCoinEntity)
	if err != nil {
		return nil, err
	}

	savedPlayerCoin.PlayerID = coinAddingReq.PlayerID

	return savedPlayerCoin.ToPlayerCoinModel(), nil
}

func (s *playerCoinServiceImpl) CoinShowing(playerID string) *playerCoinModels.PlayerCoinShowing {
	playerCoinShowing, err := s.playerCoinRepo.CoinShowing(playerID)
	if err != nil {
		return &playerCoinModels.PlayerCoinShowing{PlayerID: playerCoinShowing.PlayerID, Coin: 0}
	}

	return playerCoinShowing
}
