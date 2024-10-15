package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	_ "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
	playerCoinExceptions "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/exceptions"
	playerCoinModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/models"
	"gorm.io/gorm"
)

type playerCoinRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerCoinRepositoryImpl(db databases.Database, logger echo.Logger) PlayerCoinRepository {
	return &playerCoinRepositoryImpl{db, logger}
}

func (r *playerCoinRepositoryImpl) CoinAdding(tx *gorm.DB, playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	savedPlayerCoin := new(entities.PlayerCoin)

	if err := conn.Create(playerCoinEntity).Scan(savedPlayerCoin).Error; err != nil {
		r.logger.Errorf("Failed to add player coin: %s", err.Error())
		return nil, &playerCoinExceptions.CoinAdding{}
	}

	return savedPlayerCoin, nil
}

func (r *playerCoinRepositoryImpl) CoinShowing(playerID string) (*playerCoinModels.PlayerCoinShowing, error) {
	playerCoinShowing := new(playerCoinModels.PlayerCoinShowing)

	if err := r.db.Connect().Model(&entities.PlayerCoin{}).Where("player_id = ?", playerID).
		Select("player_id, sum(amount) as coin").Group("player_id").Scan(playerCoinShowing).Error; err != nil {
		r.logger.Errorf("Failed to show player coin")
		return nil, &playerCoinExceptions.PlayerCoinShowing{}
	}

	return playerCoinShowing, nil
}
