package routes

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/controllers"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/repositories"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/services"
)

func (r *Router) registerItemShopRoutes() {
	itemShopRoutes := r.app.Group("/v1/item-shop")

	itemShopRepo := repositories.NewItemShopRepositoryImpl(r.db, r.logger)
	itemShopService := services.NewItemShopServiceImpl(itemShopRepo)
	itemShopController := controllers.NewItemShopControllerImpl(itemShopService)

	itemShopRoutes.GET("", itemShopController.Listing)

}
