package services

import (
	adminRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/admin/repositories"
	playerRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/player/repositories"
)

type googleOAuth2Service struct {
	playerRepo playerRepositories.PlayerRepository
	adminRepo  adminRepositories.AdminRepository
}

func NewGoogleOAuth2Service(playerRepo playerRepositories.PlayerRepository, adminRepo adminRepositories.AdminRepository) Oauth2Service {
	return &googleOAuth2Service{playerRepo, adminRepo}
}
