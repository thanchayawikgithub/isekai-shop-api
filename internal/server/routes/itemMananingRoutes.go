package routes

import (
	itemManagingControllers "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/controllers"
	itemManagingRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/repositories"
	itemManagingServices "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/services"
)

func (r *Router) registerItemManagingRoutes() {
	itemManagingRoutes := r.app.Group("/v1/item-managing")

	itemManagingRepo := itemManagingRepositories.NewItemManagingRepositoryImpl(r.db, r.logger)
	itemManagingService := itemManagingServices.NewItemManagingServiceImpl(itemManagingRepo)
	itemManagingController := itemManagingControllers.NewItemManagingControllerImpl(itemManagingService)

	itemManagingRoutes.POST("", itemManagingController.Creating)
}
