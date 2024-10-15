package routes

import (
	playerCoinControllers "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/controllers"
	playerCoinRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/repositories"
	playerCoinServices "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/playerCoin/services"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/server/middlewares"
)

func (r *Router) registerPlayerCoinRoutes(authMiddleWare *middlewares.AuthMiddleWare) {
	playerCoinRoutes := r.app.Group("/v1/player-coin")

	playerCoinRepo := playerCoinRepositories.NewPlayerCoinRepositoryImpl(r.db, r.logger)
	playerCoinService := playerCoinServices.NewPlayerCoinServiceImpl(playerCoinRepo)
	playerCoinController := playerCoinControllers.NewPlayerCoinControllerImpl(playerCoinService)

	playerCoinRoutes.POST("", playerCoinController.CoinAdding, authMiddleWare.PlayerAuthorize)
	playerCoinRoutes.GET("", playerCoinController.CoinShowing, authMiddleWare.PlayerAuthorize)
}
