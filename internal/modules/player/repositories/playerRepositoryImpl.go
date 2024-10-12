package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	playerExceptions "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/player/exceptons"
)

type playerRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerRepositoryImpl(db databases.Database, logger echo.Logger) PlayerRepository {
	return &playerRepositoryImpl{db, logger}
}

func (r *playerRepositoryImpl) Creating(playerEntity *entities.Player) (*entities.Player, error) {
	savedPlayer := new(entities.Player)

	if err := r.db.Connect().Create(playerEntity).Scan(savedPlayer).Error; err != nil {
		r.logger.Errorf("Failed to creating player", err.Error())
		return nil, &playerExceptions.PlayerCreating{PlayerID: playerEntity.ID}
	}

	return savedPlayer, nil
}

func (r *playerRepositoryImpl) FindByID(playerID string) (*entities.Player, error) {
	player := new(entities.Player)

	if err := r.db.Connect().Where("id = ?", playerID).First(player).Error; err != nil {
		r.logger.Errorf("Failed to find player by ID", err.Error())
		return nil, &playerExceptions.PlayerNotFound{PlayerID: playerID}
	}

	return player, nil
}
