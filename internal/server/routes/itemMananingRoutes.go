package routes

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/controllers"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/repositories"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/services"
)

func (r *Router) registerItemManagingRoutes() {
	itemManagingRoutes := r.app.Group("/v1/item-managing")

	itemManagingRepo := repositories.NewItemManagingRepositoryImpl(r.db, r.logger)
	itemManagingService := services.NewItemManagingServiceImpl(itemManagingRepo)
	itemManagingController := controllers.NewItemManagingControllerImpl(itemManagingService)

	_ = itemManagingController
	_ = itemManagingRoutes
}
