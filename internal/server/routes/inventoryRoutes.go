package routes

import (
	inventoryControllers "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/inventory/controllers"
	inventoryRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/inventory/repositories"
	inventoryServices "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/inventory/services"
	itemShopRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/repositories"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/server/middlewares"
)

func (r *Router) registerInventoryRoutes(authMiddleWare *middlewares.AuthMiddleWare) {
	inventoryRoutes := r.app.Group("/v1/inventory")

	itemShopRepo := itemShopRepositories.NewItemShopRepositoryImpl(r.db, r.logger)
	inventoryRepo := inventoryRepositories.NewInventoryRepositoryImpl(r.db, r.logger)
	inventoryService := inventoryServices.NewInventoryServiceImpl(inventoryRepo, itemShopRepo)
	inventoryController := inventoryControllers.NewInventoryControllerImpl(inventoryService)

	inventoryRoutes.GET("", inventoryController.Listing, authMiddleWare.PlayerAuthorize)
}
