package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases"
	_ "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/exceptions"
	_ "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/models"
)

type playerCoinRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerCoinRepositoryImpl(db databases.Database, logger echo.Logger) PlayerCoinRepository {
	return &playerCoinRepositoryImpl{db, logger}
}
