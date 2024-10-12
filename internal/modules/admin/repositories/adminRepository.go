package repositories

import "github.com/thanchayawikgithub/isekai-shop-api/internal/entities"

type AdminRepository interface {
	Creating(adminEntity *entities.Admin) (*entities.Admin, error)
	FindByID(adminID string) (*entities.Admin, error)
}
