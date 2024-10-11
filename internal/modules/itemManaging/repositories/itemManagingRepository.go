package repositories

import "github.com/thanchayawikgithub/isekai-shop-api/internal/entities"

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
}
