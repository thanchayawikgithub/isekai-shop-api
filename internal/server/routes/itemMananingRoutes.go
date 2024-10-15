package routes

import (
	itemManagingControllers "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/controllers"
	itemManagingRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/repositories"
	itemManagingServices "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/services"
	itemShopRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/repositories"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/server/middlewares"
)

func (r *Router) registerItemManagingRoutes(authMiddleWare *middlewares.AuthMiddleWare) {
	itemManagingRoutes := r.app.Group("/v1/item-managing")

	itemShopRepo := itemShopRepositories.NewItemShopRepositoryImpl(r.db, r.logger)
	itemManagingRepo := itemManagingRepositories.NewItemManagingRepositoryImpl(r.db, r.logger)
	itemManagingService := itemManagingServices.NewItemManagingServiceImpl(itemManagingRepo, itemShopRepo)
	itemManagingController := itemManagingControllers.NewItemManagingControllerImpl(itemManagingService)

	itemManagingRoutes.POST("", itemManagingController.Creating, authMiddleWare.AdminAuthorize)
	itemManagingRoutes.PATCH("/:itemID", itemManagingController.Editing)
	itemManagingRoutes.DELETE("/:itemID", itemManagingController.Archiving)
}
