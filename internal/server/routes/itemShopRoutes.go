package routes

import (
	itemShopControllers "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/controllers"
	itemShopRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/repositories"
	itemShopServices "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/services"
)

func (r *Router) registerItemShopRoutes() {
	itemShopRoutes := r.app.Group("/v1/item-shop")

	itemShopRepo := itemShopRepositories.NewItemShopRepositoryImpl(r.db, r.logger)
	itemShopService := itemShopServices.NewItemShopServiceImpl(itemShopRepo)
	itemShopController := itemShopControllers.NewItemShopControllerImpl(itemShopService)

	itemShopRoutes.GET("", itemShopController.Listing)
}
