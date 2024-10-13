package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	adminExceptions "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/admin/exceptions"
)

type adminRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewAdminRepositoryImpl(db databases.Database, logger echo.Logger) AdminRepository {
	return &adminRepositoryImpl{db, logger}
}

func (r *adminRepositoryImpl) Creating(adminEntity *entities.Admin) (*entities.Admin, error) {
	savedAdmin := new(entities.Admin)

	if err := r.db.Connect().Create(adminEntity).Scan(savedAdmin).Error; err != nil {
		r.logger.Errorf("Failed to creating player", err.Error())
		return nil, &adminExceptions.AdminCreating{AdminID: adminEntity.ID}
	}

	return savedAdmin, nil
}

func (r *adminRepositoryImpl) FindByID(adminID string) (*entities.Admin, error) {
	admin := new(entities.Admin)

	if err := r.db.Connect().Where("id = ?", adminID).First(admin).Error; err != nil {
		r.logger.Errorf("Failed to find admin by ID", err.Error())
		return nil, &adminExceptions.AdminNotFound{AdminID: adminID}
	}

	return admin, nil
}
