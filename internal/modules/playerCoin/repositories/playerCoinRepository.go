package repositories

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	playerCoinModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/models"
)

type PlayerCoinRepository interface {
	CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	CoinShowing(playerID string) (*playerCoinModels.PlayerCoinShowing, error)
}
