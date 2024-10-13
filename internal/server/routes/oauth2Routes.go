package routes

import (
	adminRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/admin/repositories"
	oauth2Controllers "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/oauth2/controllers"
	oauth2Services "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/oauth2/services"
	playerRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/player/repositories"
)

func (r *Router) registerOAuth2Routes() {
	oauth2Routes := r.app.Group("/v1/oauth2/google")

	playerRepo := playerRepositories.NewPlayerRepositoryImpl(r.db, r.logger)
	adminRepo := adminRepositories.NewAdminRepositoryImpl(r.db, r.logger)

	oauth2Service := oauth2Services.NewGoogleOAuth2Service(playerRepo, adminRepo)
	oauth2Controller := oauth2Controllers.NewGoogleOAuth2Controller(oauth2Service, r.config.OAuth2, r.app.Logger)

	oauth2Routes.GET("/player/login", oauth2Controller.PlayerLogin)
	oauth2Routes.GET("/admin/login", oauth2Controller.AdminLogin)
	oauth2Routes.GET("/player/login/callback", oauth2Controller.PlayerLoginCallback)
	oauth2Routes.GET("/admin/login/callback", oauth2Controller.AdminLoginCallback)
	oauth2Routes.DELETE("/logout", oauth2Controller.Logout)
}
