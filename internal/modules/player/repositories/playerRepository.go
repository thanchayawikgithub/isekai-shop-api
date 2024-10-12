package repositories

import "github.com/thanchayawikgithub/isekai-shop-api/internal/entities"

type PlayerRepository interface {
	Creating(playerEntity *entities.Player) (*entities.Player, error)
	FindByID(playerID string) (*entities.Player, error)
}
