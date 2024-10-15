package routes

import (
	inventoryRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/inventory/repositories"
	itemShopControllers "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/controllers"
	itemShopRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/repositories"
	itemShopServices "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/services"
	playerCoinRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/repositories"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/server/middlewares"
)

func (r *Router) registerItemShopRoutes(authMiddleWare *middlewares.AuthMiddleWare) {
	itemShopRoutes := r.app.Group("/v1/item-shop")

	playerCoinRepo := playerCoinRepositories.NewPlayerCoinRepositoryImpl(r.db, r.logger)
	inventoryRepo := inventoryRepositories.NewInventoryRepositoryImpl(r.db, r.logger)
	itemShopRepo := itemShopRepositories.NewItemShopRepositoryImpl(r.db, r.logger)
	itemShopService := itemShopServices.NewItemShopServiceImpl(itemShopRepo, playerCoinRepo, inventoryRepo)
	itemShopController := itemShopControllers.NewItemShopControllerImpl(itemShopService)

	itemShopRoutes.GET("", itemShopController.Listing)
	itemShopRoutes.POST("/buying", itemShopController.Buying, authMiddleWare.PlayerAuthorize)
}
