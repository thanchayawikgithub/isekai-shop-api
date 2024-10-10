package server

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/controllers"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/repositories"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/services"
)

func (s *echoServer) registerItemShopRouter() {
	itemShop := s.app.Group("/v1/item-shop")

	itemShopRepo := repositories.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := services.NewItemShopServiceImpl(itemShopRepo)
	itemShopController := controllers.NewItemShopControllerImpl(itemShopService)

	itemShop.GET("", itemShopController.Listing)

}
