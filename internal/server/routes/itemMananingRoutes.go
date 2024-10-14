package routes

import (
	adminRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/admin/repositories"
	itemManagingControllers "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/controllers"
	itemManagingRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/repositories"
	itemManagingServices "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemManaging/services"
	itemShopRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/itemShop/repositories"
	oauth2Controllers "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/oauth2/controllers"
	oauth2Services "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/oauth2/services"
	playerRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/player/repositories"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/server/middlewares"
)

func (r *Router) registerItemManagingRoutes() {
	itemManagingRoutes := r.app.Group("/v1/item-managing")

	itemShopRepo := itemShopRepositories.NewItemShopRepositoryImpl(r.db, r.logger)
	itemManagingRepo := itemManagingRepositories.NewItemManagingRepositoryImpl(r.db, r.logger)
	itemManagingService := itemManagingServices.NewItemManagingServiceImpl(itemManagingRepo, itemShopRepo)
	itemManagingController := itemManagingControllers.NewItemManagingControllerImpl(itemManagingService)

	playerRepo := playerRepositories.NewPlayerRepositoryImpl(r.db, r.logger)
	adminRepo := adminRepositories.NewAdminRepositoryImpl(r.db, r.logger)
	oauth2Service := oauth2Services.NewGoogleOAuth2Service(playerRepo, adminRepo)
	oauth2Controller := oauth2Controllers.NewGoogleOAuth2Controller(oauth2Service, r.config.OAuth2, r.app.Logger)
	authMiddleWare := middlewares.NewAuthMiddleware(oauth2Controller, r.config.OAuth2, r.logger)

	itemManagingRoutes.POST("", itemManagingController.Creating, authMiddleWare.AdminAuthorize)
	itemManagingRoutes.PATCH("/:itemID", itemManagingController.Editing)
	itemManagingRoutes.DELETE("/:itemID", itemManagingController.Archiving)
}
